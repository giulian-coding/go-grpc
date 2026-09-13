package server

import (
	"context"

	pb "github.com/giulian-coding/hosts/internal/grpc"
)

type Server struct {
	pb.UnimplementedJobServiceServer
}

func (s *Server) SendJob(ctx context.Context, in *pb.JobRequest) (*pb.JobReply, error) {
	return &pb.JobReply{Message: "Received" + in.Id}, nil
}
