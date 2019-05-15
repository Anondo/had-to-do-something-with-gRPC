package client

import (
	"context"
	pb "emprpc/employee"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// MakeRequest makes a gRPC request to the gRPC server
func MakeRequest() error {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	clnt := pb.NewEmployeeServiceClient(conn)
	id := viper.GetInt("id")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	resp, err := clnt.GetEmployee(ctx, &pb.EmployeeRequest{Id: int32(id)})

	if err != nil {
		spew.Dump("Here")
		return err
	}
	log.Println(resp)
	return nil
}
