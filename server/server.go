package server

import (
	"context"
	"emprpc/db"
	pb "emprpc/employee"
	"encoding/json"
	"net"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-redis/redis"
	"google.golang.org/grpc"
)

type empServer struct{}

func (e *empServer) GetEmployee(ctx context.Context, r *pb.EmployeeRequest) (*pb.Employee, error) {
	rds := db.GetDB()
	bb, err := rds.Get(strconv.Itoa(int(r.Id))).Result()
	if err != nil {
		if err != redis.Nil {
			return nil, err
		}
	}
	spew.Dump(bb)
	emp := pb.Employee{}

	if err := json.Unmarshal([]byte(bb), &emp); err != nil {
		return nil, err
	}

	return &emp, nil
}

// StartServer starts the gRPC server
func StartServer() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}
	srvr := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(srvr, &empServer{})

	return srvr.Serve(lis)
}
