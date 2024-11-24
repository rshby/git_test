package service

import (
	"context"
	"git_test/proto/pb"
	"github.com/sirupsen/logrus"
)

type studentService struct {
	pb.UnimplementedStudentServiceServer
}

// NewStudentService is function to create new instance of studentService
func NewStudentService() pb.StudentServiceServer {
	return &studentService{}
}

func (s *studentService) RegisterStudent(ctx context.Context, request *pb.CreateStudentRequestDTO) (*pb.CreateStudentResponseDTO, error) {
	logrus.Infof("receive request new student : [%s]", request.String())
	return &pb.CreateStudentResponseDTO{Message: "success create new student"}, nil
}

func (s *studentService) GetStudentByID(ctx context.Context, request *pb.GetStudentByIDRequestDTO) (*pb.Student, error) {
	//TODO implement me
	panic("implement me")
}
