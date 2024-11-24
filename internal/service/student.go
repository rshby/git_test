package service

import (
	"context"
	"fmt"
	"git_test/proto/pb"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/semaphore"
	"google.golang.org/grpc"
	"sync"
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

// GetStudentBySchoolID is method to send with stream
func (s *studentService) GetStudentBySchoolID(request *pb.GetStudentBySchoolIDRequestDTO, srv grpc.ServerStreamingServer[pb.Student]) error {
	var (
		wg       = &sync.WaitGroup{}
		semLimit = semaphore.NewWeighted(100)
	)

	for i := 0; i < 100; i++ {
		if err := semLimit.Acquire(context.Background(), 1); err != nil {
			continue
		}

		wg.Add(1)
		go func(wg *sync.WaitGroup, semLimit *semaphore.Weighted, id int) {
			defer func() {
				semLimit.Release(1)
				wg.Done()
			}()

			if err := srv.Send(&pb.Student{
				Id:      int32(id + 1),
				Name:    "Reo Sahobby",
				Email:   "reoshby@gmail.com",
				Phone:   "083863890419",
				Address: "Jakarta Selatan",
				School:  fmt.Sprintf("%d", request.GetSchoolID()),
			}); err != nil {
				return
			}
		}(wg, semLimit, i)
	}

	wg.Wait()
	return nil
}
