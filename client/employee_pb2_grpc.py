# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import employee_pb2 as employee__pb2


class EmployeeServiceStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetEmployee = channel.unary_unary(
        '/EmployeeService/GetEmployee',
        request_serializer=employee__pb2.EmployeeRequest.SerializeToString,
        response_deserializer=employee__pb2.Employee.FromString,
        )
    self.GetEmployeeList = channel.unary_stream(
        '/EmployeeService/GetEmployeeList',
        request_serializer=employee__pb2.Nothing.SerializeToString,
        response_deserializer=employee__pb2.Employee.FromString,
        )
    self.SetEmployee = channel.unary_unary(
        '/EmployeeService/SetEmployee',
        request_serializer=employee__pb2.Employee.SerializeToString,
        response_deserializer=employee__pb2.ResponseMessage.FromString,
        )
    self.DeleteEmployee = channel.unary_unary(
        '/EmployeeService/DeleteEmployee',
        request_serializer=employee__pb2.EmployeeRequest.SerializeToString,
        response_deserializer=employee__pb2.ResponseMessage.FromString,
        )
    self.UpdateEmployee = channel.unary_unary(
        '/EmployeeService/UpdateEmployee',
        request_serializer=employee__pb2.EmployeeUpdateRequest.SerializeToString,
        response_deserializer=employee__pb2.ResponseMessage.FromString,
        )


class EmployeeServiceServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def GetEmployee(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetEmployeeList(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SetEmployee(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DeleteEmployee(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def UpdateEmployee(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_EmployeeServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetEmployee': grpc.unary_unary_rpc_method_handler(
          servicer.GetEmployee,
          request_deserializer=employee__pb2.EmployeeRequest.FromString,
          response_serializer=employee__pb2.Employee.SerializeToString,
      ),
      'GetEmployeeList': grpc.unary_stream_rpc_method_handler(
          servicer.GetEmployeeList,
          request_deserializer=employee__pb2.Nothing.FromString,
          response_serializer=employee__pb2.Employee.SerializeToString,
      ),
      'SetEmployee': grpc.unary_unary_rpc_method_handler(
          servicer.SetEmployee,
          request_deserializer=employee__pb2.Employee.FromString,
          response_serializer=employee__pb2.ResponseMessage.SerializeToString,
      ),
      'DeleteEmployee': grpc.unary_unary_rpc_method_handler(
          servicer.DeleteEmployee,
          request_deserializer=employee__pb2.EmployeeRequest.FromString,
          response_serializer=employee__pb2.ResponseMessage.SerializeToString,
      ),
      'UpdateEmployee': grpc.unary_unary_rpc_method_handler(
          servicer.UpdateEmployee,
          request_deserializer=employee__pb2.EmployeeUpdateRequest.FromString,
          response_serializer=employee__pb2.ResponseMessage.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'EmployeeService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))


class DeptServiceStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetDept = channel.unary_unary(
        '/DeptService/GetDept',
        request_serializer=employee__pb2.DeptRequest.SerializeToString,
        response_deserializer=employee__pb2.Dept.FromString,
        )
    self.GetDeptList = channel.unary_stream(
        '/DeptService/GetDeptList',
        request_serializer=employee__pb2.Nothing.SerializeToString,
        response_deserializer=employee__pb2.Dept.FromString,
        )
    self.SetDept = channel.unary_unary(
        '/DeptService/SetDept',
        request_serializer=employee__pb2.Dept.SerializeToString,
        response_deserializer=employee__pb2.ResponseMessage.FromString,
        )
    self.DeleteDepartment = channel.unary_unary(
        '/DeptService/DeleteDepartment',
        request_serializer=employee__pb2.DeptRequest.SerializeToString,
        response_deserializer=employee__pb2.ResponseMessage.FromString,
        )
    self.UpdateDepartment = channel.unary_unary(
        '/DeptService/UpdateDepartment',
        request_serializer=employee__pb2.DeptUpdateRequest.SerializeToString,
        response_deserializer=employee__pb2.ResponseMessage.FromString,
        )


class DeptServiceServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def GetDept(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetDeptList(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SetDept(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DeleteDepartment(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def UpdateDepartment(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_DeptServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetDept': grpc.unary_unary_rpc_method_handler(
          servicer.GetDept,
          request_deserializer=employee__pb2.DeptRequest.FromString,
          response_serializer=employee__pb2.Dept.SerializeToString,
      ),
      'GetDeptList': grpc.unary_stream_rpc_method_handler(
          servicer.GetDeptList,
          request_deserializer=employee__pb2.Nothing.FromString,
          response_serializer=employee__pb2.Dept.SerializeToString,
      ),
      'SetDept': grpc.unary_unary_rpc_method_handler(
          servicer.SetDept,
          request_deserializer=employee__pb2.Dept.FromString,
          response_serializer=employee__pb2.ResponseMessage.SerializeToString,
      ),
      'DeleteDepartment': grpc.unary_unary_rpc_method_handler(
          servicer.DeleteDepartment,
          request_deserializer=employee__pb2.DeptRequest.FromString,
          response_serializer=employee__pb2.ResponseMessage.SerializeToString,
      ),
      'UpdateDepartment': grpc.unary_unary_rpc_method_handler(
          servicer.UpdateDepartment,
          request_deserializer=employee__pb2.DeptUpdateRequest.FromString,
          response_serializer=employee__pb2.ResponseMessage.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'DeptService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
