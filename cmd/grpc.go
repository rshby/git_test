package main

import (
	"context"
	"git_test/interceptor"
	"git_test/internal/service"
	"git_test/proto/pb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logrus.Fatal(err)
	}

	// create new gRpc server
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(interceptor.WithLogInterceptor(), interceptor.WithErrorInterceptor()),
	)

	// create new instance service
	customerService := service.NewCustomerService()
	studentService := service.NewStudentService()

	// register service
	pb.RegisterCustomerServiceServer(srv, customerService)
	pb.RegisterStudentServiceServer(srv, studentService)

	reflection.Register(srv)

	// create gRPC Gateway Server
	var gwServer *http.Server

	chanSignal := make(chan os.Signal, 1)
	chanErr := make(chan error, 1)
	chanQuit := make(chan struct{}, 1)

	signal.Notify(chanSignal, os.Interrupt)

	// spawn goroutine : idle to consume signal if any interrupt or error
	go func() {
		for {
			select {
			case <-chanSignal:
				logrus.Info("receive interrupt signal ⚠️")
				gracefullShutdown(gwServer, srv)
				chanQuit <- struct{}{}
				return
			case e := <-chanErr:
				logrus.Infof("receive error chan ⚠️ : %s", e.Error())
				gracefullShutdown(gwServer, srv)
				chanQuit <- struct{}{}
				return
			}
		}
	}()

	// spawn goroutine : run gRPC Server and gRPC Gateway Server
	go func() {
		// spawn goroutine : run gRPC server
		go func() {
			logrus.Infof("running gRPC Server Listening on : [localhost:%d] ✅", 50051)
			if err = srv.Serve(lis); err != nil {
				chanErr <- err
				return
			}
		}()

		// spawn goroutine : run gRPC Gateway Server
		go func() {
			var (
				muxServer  = runtime.NewServeMux()
				conn       *grpc.ClientConn
				err        error
				maxRetries = 3
			)

			for i := 1; i <= maxRetries; i++ {
				conn, err = grpc.Dial("localhost:50051", grpc.WithInsecure())
				if err == nil {
					logrus.Info("gRPC Gateway success dial to gRPC Server ✅")
					break
				}

				time.Sleep(2 * time.Second)
			}

			if err != nil {
				chanErr <- err
				return
			}

			if err = pb.RegisterCustomerServiceHandler(context.Background(), muxServer, conn); err != nil {
				chanErr <- err
				return
			}

			gwServer = &http.Server{
				Addr:    ":4002",
				Handler: muxServer,
			}

			logrus.Info("running gRPC Gateway Listening on [localhost:4002] ✅")
			if err = gwServer.ListenAndServe(); err != nil {
				chanErr <- err
				return
			}
		}()
	}()

	_ = <-chanQuit
	close(chanQuit)
	close(chanSignal)
	close(chanErr)

	logrus.Info("server exit ‼️")
}

// gracefullShutdown is function to close server gracefully
func gracefullShutdown(httpServer *http.Server, grpcServer *grpc.Server) {
	var wg = &sync.WaitGroup{}

	// spawn goroutine : close gRPC Gateway Server gracefully
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		if httpServer != nil {
			ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancelFunc()
			if err := httpServer.Shutdown(ctx); err != nil {
				httpServer.Close()
				logrus.Info("force close http Server 🛑")
				return
			}
			logrus.Info("success shutdown gRPC Gateway Server gracefully 🛑")

			if err := httpServer.Close(); err != nil {
				logrus.Infof("cant close http Server 🛑 : %s", err.Error())
				return
			}
			logrus.Info("success stop gRPC Gateway Server 🛑")
		}
	}(wg)

	// spawn goroutine : close gRPC Server gracefully
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		if grpcServer != nil {
			grpcServer.GracefulStop()
			logrus.Info("success shutdown gRPC Server gracefully 🛑")

			grpcServer.Stop()
			logrus.Info("success stop gRPC Server 🛑")
		}

	}(wg)

	wg.Wait()
}
