package main

import (
	"context"
	"fmt"
	"git_test/proto/pb"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/semaphore"
	"google.golang.org/grpc"
	"log"
	"sync"
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

	var (
		wg       = &sync.WaitGroup{}
		semLimit = semaphore.NewWeighted(int64(100000))
	)

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, semLimit *semaphore.Weighted) {
			defer wg.Done()

			if err := semLimit.Acquire(context.Background(), 1); err != nil {
				return
			}
			defer semLimit.Release(1)

			// create new student
			_, _ = externalService.StudentService.RegisterStudent(ctx, &pb.CreateStudentRequestDTO{
				Name:    "Reo Sahobby",
				Email:   "reoshby@gmail.com",
				Phone:   "083863890419",
				Address: "Jakarta Selatan",
				School:  "UPN Veteran",
			})
		}(wg, semLimit)
	}

	wg.Wait()
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
