package main

import (
	"context"
	"fmt"
	"had-to-do-something-with-gRPC/server/db"
	pb "had-to-do-something-with-gRPC/server/employee"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"gopkg.in/mgo.v2/bson"
)

const (
	port = 50051
)

type (
	empServer  struct{}
	deptServer struct{}
)

func init() {
	db.Init()
}

func (e *empServer) GetEmployee(ctx context.Context, r *pb.EmployeeRequest) (*pb.Employee, error) {
	mdb := db.GetDB()

	emp := pb.Employee{}

	if err := mdb.Collection("employee").FindOne(ctx, bson.M{"id": r.GetId()}).Decode(&emp); err != nil {
		return nil, err
	}

	return &emp, nil
}

func (e *empServer) GetEmployeeList(_ *pb.Nothing, stream pb.EmployeeService_GetEmployeeListServer) error {
	mdb := db.GetDB()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	cur, err := mdb.Collection("employee").Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var emp pb.Employee
		if err := cur.Decode(&emp); err != nil {
			return err
		}
		if err := stream.Send(&emp); err != nil {
			return err
		}
	}
	return nil
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

func (e *empServer) DeleteEmployee(ctx context.Context, r *pb.EmployeeRequest) (*pb.ResponseMessage, error) {

	mdb := db.GetDB()
	dr, err := mdb.Collection("employee").DeleteOne(ctx, bson.M{"id": r.GetId()})
	if err != nil {
		return &pb.ResponseMessage{
			Msg: "Couldnt delete employee: " + err.Error(),
		}, err
	}
	if dr.DeletedCount < 1 {
		return &pb.ResponseMessage{
			Msg: "Employee Not Found",
		}, nil
	}

	return &pb.ResponseMessage{
		Msg: "Employee deleted successfully",
	}, nil
}

func (e *empServer) UpdateEmployee(ctx context.Context, r *pb.EmployeeUpdateRequest) (*pb.ResponseMessage, error) {
	mdb := db.GetDB()

	ur, err := mdb.Collection("employee").UpdateOne(ctx, bson.M{"id": r.GetId()}, r.GetEmp())
	if err != nil {
		return &pb.ResponseMessage{
			Msg: "Could not update employee: " + err.Error(),
		}, err
	}

	if ur.MatchedCount < 1 || ur.ModifiedCount < 1 {
		return &pb.ResponseMessage{
			Msg: "No employee documents was matched or modified",
		}, nil
	}

	return &pb.ResponseMessage{
		Msg: "Employee information updated",
	}, nil
}

func (d *deptServer) GetDept(ctx context.Context, r *pb.DeptRequest) (*pb.Dept, error) {
	mdb := db.GetDB()

	dept := pb.Dept{}

	if err := mdb.Collection("department").FindOne(ctx, bson.M{"id": r.GetId()}).Decode(&dept); err != nil {
		return nil, err
	}

	return &dept, nil
}

func (d *deptServer) GetDeptList(_ *pb.Nothing, stream pb.DeptService_GetDeptListServer) error {
	mdb := db.GetDB()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	cur, err := mdb.Collection("department").Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var dpt pb.Dept
		if err := cur.Decode(&dpt); err != nil {
			return err
		}
		if err := stream.Send(&dpt); err != nil {
			return err
		}
	}
	return nil
}

func (d *deptServer) SetDept(ctx context.Context, dept *pb.Dept) (*pb.ResponseMessage, error) {
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

func (d *deptServer) DeleteDepartment(ctx context.Context, r *pb.DeptRequest) (*pb.ResponseMessage, error) {

	mdb := db.GetDB()

	dr, err := mdb.Collection("department").DeleteOne(ctx, bson.M{"id": r.GetId()})

	if err != nil {
		return &pb.ResponseMessage{
			Msg: "Couldnt delete department: " + err.Error(),
		}, err
	}

	if dr.DeletedCount < 1 {
		return &pb.ResponseMessage{
			Msg: "Department not found",
		}, nil
	}
	return &pb.ResponseMessage{
		Msg: "Department deleted successfully",
	}, nil
}

func (d *deptServer) UpdateDepartment(ctx context.Context, r *pb.DeptUpdateRequest) (*pb.ResponseMessage, error) {
	mdb := db.GetDB()

	ur, err := mdb.Collection("department").UpdateOne(ctx, bson.M{"id": r.GetId()}, r.GetDepartment())
	if err != nil {
		return &pb.ResponseMessage{
			Msg: "could not update department document: " + err.Error(),
		}, err
	}

	if ur.MatchedCount < 1 || ur.ModifiedCount < 1 {
		return &pb.ResponseMessage{
			Msg: "No department document found or modified",
		}, nil
	}

	return &pb.ResponseMessage{
		Msg: "Department information updated",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err.Error())
	}
	srvr := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(srvr, &empServer{})
	pb.RegisterDeptServiceServer(srvr, &deptServer{})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT)

	log.Printf("gRPC server running  at localhost:%d ...\n", port)
	go func() {
		if err := srvr.Serve(lis); err != nil {
			log.Fatal(err.Error())
		}
	}()

	<-stop

	fmt.Println("Bye bye")

}
