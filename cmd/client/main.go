package main

import (
	"context"
	"log"
	"time"
	"uuid"

	pb "github.com/giulian-coding/hosts/internal/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	id := uuid.NewV7()

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewJobServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendJob(ctx, &pb.JobRequest{
		Id: id.String(),
		Payload: &pb.Payload{
			Task:       "dosomething",
			Parameters: "someparameters",
		},
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
