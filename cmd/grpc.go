package main

import (
	"git_test/interceptor"
	"git_test/internal/service"
	"git_test/proto/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
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
	if err = srv.Serve(lis); err != nil {
		logrus.Fatal(err)
	}

	srv.Stop()
	logrus.Info("server exit ‼️")
}
