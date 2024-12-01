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

	// run server
	logrus.Infof("running grpc server on : [localhost:%d] ✅", 50051)
	go func() {
		if err = srv.Serve(lis); err != nil {
			logrus.Fatal(err)
		}
	}()

	// run grpc gateway
	muxServer := runtime.NewServeMux()
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		logrus.Fatal(err)
	}

	if err = pb.RegisterCustomerServiceHandler(context.Background(), muxServer, conn); err != nil {
		logrus.Fatal(err)
	}

	gwServer := &http.Server{
		Addr:    ":4002",
		Handler: muxServer,
	}

	logrus.Info("running gRPC Gateway Listening on [localhost:4002] ✅")
	if err = gwServer.ListenAndServe(); err != nil {
		logrus.Fatal(err)
	}

	srv.Stop()
	logrus.Info("server exit ‼️")
}
