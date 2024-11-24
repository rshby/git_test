package main

import (
	"context"
	"fmt"
	"git_test/proto/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn := InitExternalServiceGRPC()
	defer conn.Close()

	ctx := context.Background()

	// get customer, hit from grpc
	customer, err := externalService.CustomerService.GetCustomerByEmail(ctx, &pb.GetCustomerByEmailRequest{Email: string("reoshby@gmail.com")})
	if err != nil {
		logrus.Error(err)
	}

	fmt.Println(customer)

	// create new student
	resultRegister, err := externalService.StudentService.RegisterStudent(ctx, &pb.CreateStudentRequestDTO{
		Name:    "Reo Sahobby",
		Email:   "reoshby@gmail.com",
		Phone:   "083863890419",
		Address: "Jakarta Selatan",
		School:  "UPN Veteran",
	})
	if err != nil {
		logrus.Error(err)
	}

	fmt.Println(resultRegister)
}

var externalService ExternalService

type ExternalService struct {
	CustomerService pb.CustomerServiceClient
	StudentService  pb.StudentServiceClient
}

func InitExternalServiceGRPC() *grpc.ClientConn {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to service: %v", err)
		return nil
	}
	
	externalService = ExternalService{
		CustomerService: pb.NewCustomerServiceClient(conn),
		StudentService:  pb.NewStudentServiceClient(conn),
	}

	return conn
}
