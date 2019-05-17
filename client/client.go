package client

import (
	"context"
	pb "emprpc/employee"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

var (
	empClient  pb.EmployeeServiceClient
	deptClient pb.DeptServiceClient
)

const (
	host = "localhost"
	port = 50051
)

func init() {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	empClient = pb.NewEmployeeServiceClient(conn)
	deptClient = pb.NewDeptServiceClient(conn)
}

// FetchDeptRequest fetches a  departement data
func FetchDeptRequest(id int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	dept, err := deptClient.GetDept(ctx, &pb.DeptRequest{Id: id})
	if err != nil {
		return err
	}
	log.Println(dept)
	return nil
}

// SetDeptRequest creates a new deptartment
func SetDeptRequest() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	resp, err := deptClient.SetDept(ctx, &pb.Dept{
		Name: "Tech",
	})
	log.Println(resp.GetMsg())

	return err

}

// FetchEmpRequest fetches a  employee data
func FetchEmpRequest(id int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	emp, err := empClient.GetEmployee(ctx, &pb.EmployeeRequest{Id: id})
	if err != nil {
		return err
	}
	log.Println(emp)
	return nil
}

// SetEmpRequest creates a new employee
func SetEmpRequest() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	dept, err := deptClient.GetDept(ctx, &pb.DeptRequest{Id: 1})
	if err != nil {
		return err
	}

	resp, err := empClient.SetEmployee(ctx, &pb.Employee{
		Name:        "ABCD",
		Designation: "Software Engineer",
		Salary:      120000,
		Department:  dept,
	})
	log.Println(resp.GetMsg())

	return err

}
