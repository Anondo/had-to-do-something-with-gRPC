## Why

With http being the most common approach for building APIs, [gRPC](https://grpc.io/) is getting famous. So had to do something
with gRPC. Would have felt better if this was something more useful. Thats a thing for the future i guess.<br>

The ```server``` is written in golang & is utterly a simple **Employee Information System**. The ```rpc``` procedures are mostly
to create , read , update & delete an employee document & also their departments.<br>

The ```client``` is written in python & consumes all the api the  ```server``` offers just for demonstration.


## Running

1. Preparing source files from the proto

```
$ ./build_from_proto
```

2. Start the docker containers, in our case, the mongodb   

```
$ docker-compose up -d

```

3. inside the server directory, build the go executable binary

```
$ go build .

```

4.    Run the server

```
$ ./server

```

5. From another terminal , inside the client directory, run the python client

```

$ python client.py

```

## Things to have

1. grpc tools for both [python](https://grpc.io/docs/quickstart/python/) & [golang](https://grpc.io/docs/quickstart/go/)

1. Docker
