package server

import (
	"context"
	"emprpc/db"
	pb "emprpc/employee"
	"net"

	"google.golang.org/grpc"
	"gopkg.in/mgo.v2/bson"
)

type empServer struct{}

func (e *empServer) GetEmployee(ctx context.Context, r *pb.EmployeeRequest) (*pb.Employee, error) {
	mgo := db.GetDB()

	emp := pb.Employee{}

	if err := mgo.GetEmployee(bson.M{"id": r.GetId()}, &emp); err != nil {
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
