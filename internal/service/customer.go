package service

import (
	"context"
	pb "git_test/proto/pb"
)

type customerService struct {
	pb.UnimplementedCustomerServiceServer
}

// NewCustomerService is method to create instance of customerService
func NewCustomerService() pb.CustomerServiceServer {
	return &customerService{}
}

func (c *customerService) GetCustomerByEmail(ctx context.Context, request *pb.GetCustomerByEmailRequest) (*pb.GetCustomerByEmailResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customerService) CreateCustomer(ctx context.Context, request *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	//TODO implement me
	panic("implement me")
}
