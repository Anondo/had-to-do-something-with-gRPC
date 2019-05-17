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
type deptServer struct{}

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

func (e *deptServer) GetDept(ctx context.Context, r *pb.DeptRequest) (*pb.Dept, error) {
	mdb := db.GetDB()

	dept := pb.Dept{}

	if err := mdb.Collection("department").FindOne(ctx, bson.M{"id": r.GetId()}).Decode(&dept); err != nil {
		return nil, err
	}

	return &dept, nil
}
func (e *deptServer) SetDept(ctx context.Context, dept *pb.Dept) (*pb.ResponseMessage, error) {
	rm := &pb.ResponseMessage{}
	mdb := db.GetDB()

	id, err := mdb.Collection("department").CountDocuments(ctx, bson.M{})
	if err != nil {
		rm.Msg = "Could not count the number of documents in the collection department: " + err.Error()
		return rm, err
	}

	dept.Id = int32(id + 1)

	if _, err := mdb.Collection("department").InsertOne(ctx, dept); err != nil {
		rm.Msg = "Could not create department document: " + err.Error()
		return rm, err
	}

	rm.Msg = "Successfully created department document"
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
	pb.RegisterDeptServiceServer(srvr, &deptServer{})

	return srvr.Serve(lis)
}
