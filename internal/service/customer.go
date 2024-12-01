package service

import (
	"context"
	pb "git_test/proto/pb"
	"github.com/sirupsen/logrus"
)

type customerService struct {
	pb.UnimplementedCustomerServiceServer
}

// NewCustomerService is method to create instance of customerService
func NewCustomerService() pb.CustomerServiceServer {
	return &customerService{}
}

func (c *customerService) GetCustomerByEmail(ctx context.Context, request *pb.GetCustomerByEmailRequest) (*pb.GetCustomerByEmailResponse, error) {
	logrus.Infof("receive email from request : [%s]", request.GetEmail())

	if request.GetEmail() == "error" {
		return nil, ErrCustomerNotFound
	}

	return &pb.GetCustomerByEmailResponse{
		Customer: &pb.Customer{
			Id:      1,
			Name:    "Reo Sahobby",
			Email:   request.Email,
			Phone:   "083863890419",
			Address: "Jakarta Selatan",
		},
	}, nil
}

func (c *customerService) CreateCustomer(ctx context.Context, request *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	//TODO implement me
	panic("implement me")
}

// GetCustomerByID is method to get customer by id
func (c *customerService) GetCustomerByID(ctx context.Context, request *pb.GetCustomerByIDRequest) (*pb.Customer, error) {
	return &pb.Customer{
		Id:      request.GetId(),
		Name:    "Reo",
		Email:   "reoshby@gmail.com",
		Phone:   "081",
		Address: "Jakarta Selatan",
	}, nil
}
