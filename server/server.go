package server

import (
	"context"
	"emprpc/db"
	pb "emprpc/employee"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"gopkg.in/mgo.v2/bson"
)

const (
	port = 50051
)

type empServer struct{}

func (e *empServer) GetEmployee(ctx context.Context, r *pb.EmployeeRequest) (*pb.Employee, error) {
	mdb := db.GetDB()

	emp := pb.Employee{}

	if err := mdb.Collection("employee").FindOne(ctx, bson.M{"id": r.GetId()}).Decode(&emp); err != nil {
		return nil, err
	}

	return &emp, nil
}
func (e *empServer) SetEmployee(ctx context.Context, emp *pb.Employee) (*pb.ResponseMessage, error) {
	rm := &pb.ResponseMessage{}
	mdb := db.GetDB()

	id, err := mdb.Collection("employee").CountDocuments(ctx, bson.M{})
	if err != nil {
		rm.Msg = "Could not count the number of documents in the collection employee: " + err.Error()
		return rm, err
	}

	emp.Id = int32(id + 1)

	if _, err := mdb.Collection("employee").InsertOne(ctx, emp); err != nil {
		rm.Msg = "Could not create employee document: " + err.Error()
		return rm, err
	}

	rm.Msg = "Successfully created employee document"
	return rm, nil

}

// StartServer starts the gRPC server
func StartServer() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	srvr := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(srvr, &empServer{})

	return srvr.Serve(lis)
}
