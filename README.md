# Had To Do Something With gRPC

This branch is where both the server & client belongs to the same golang app

## Run it

```
docker-compose up -d && go build .
```

```
  ./build_from_proto
```

#### Terminal 1

Start the gRPC server

```
 ./emprpc server

```

#### Terminal 2

Create a department

```
  ./emprpc client dept set
```

Get a department

```
  ./emprpc client dept get --id=1
```

Create an employee

```
  ./emprpc client emp set
```

Get an employee

```
  ./emprpc client emp get --id=1 
```
