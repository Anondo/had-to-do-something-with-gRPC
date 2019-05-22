import grpc
import employee_pb2 as pb
import employee_pb2_grpc as pbg



def run():
    with grpc.insecure_channel("localhost:50051") as channel:
        dept_stub = pbg.DeptServiceStub(channel)
        emp_stub = pbg.EmployeeServiceStub(channel)

        print("Creating department...")
        resp = dept_stub.SetDept(pb.Dept(name = "Tech"))
        print(resp.msg)

        print("The created department: ")
        dept = dept_stub.GetDept(pb.DeptRequest(id = 1))
        print(dept)

        print("All the departments: ")
        for dept in dept_stub.GetDeptList(pb.Nothing()):
            print(dept)

        print("Creating employee...")
        resp = emp_stub.SetEmployee(pb.Employee(name="ABCD",designation="Software Engineer",salary=120000,
                                                department=dept))
        print(resp.msg)

        print("The employee created: ")
        employee = emp_stub.GetEmployee(pb.EmployeeRequest(id = 1))
        print(employee)

        print("All the employees:")
        for emp in emp_stub.GetEmployeeList(pb.Nothing()):
            print(emp)

        print("Updating the department: ")
        dept.name =  "Technology"
        print(dept_stub.UpdateDepartment(pb.DeptUpdateRequest(id = 1 , department = dept)).msg)

        print("Updating the employee: ")
        employee.salary = 2200000
        print(emp_stub.UpdateEmployee(pb.EmployeeUpdateRequest(id = 1 , emp = employee)).msg)

        print("The updated department: ")
        print(dept_stub.GetDept(pb.DeptRequest(id = 1)))

        print("The updated employee: ")
        print(emp_stub.GetEmployee(pb.EmployeeRequest(id = 1)))

        print("Deleting the employee: ")
        print(emp_stub.DeleteEmployee(pb.EmployeeRequest(id = 1)).msg)

        print("Deleting the department: ")
        print(dept_stub.DeleteDepartment(pb.DeptRequest(id = 1)).msg)

        print("All the employees:")
        for emp in emp_stub.GetEmployeeList(pb.Nothing()):
            print(emp)

        print("All the departments: ")
        for dept in dept_stub.GetDeptList(pb.Nothing()):
            print(dept)



if __name__ == '__main__':
    run()
