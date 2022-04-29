## Compile proto:
Under the root of this project `grpc-hello`
```shell
protoc --go_out=. --go-grpc_out=. proto/hello.proto
```
or
```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/hello.proto
```

## Build and Run:

### Option 1
+ open a terminal, starting server
```shell
cd server
go build && ./server
```
+ open another terminal, starting client
```shell
cd client
go build && ./client
```

### Option 2
+ open a terminal, starting server
```shell
cd server
go run server.go
```
+ open another terminal, starting client
```shell
cd client
go run client.go
```

## make a docker image and run
+ write a [docker file](./Dockerfile) used for building a docker image

+ build a docker image
```shell
docker build -t chinlying/grpc-hello:V1.0 .
```

+ put the generated image into a container and run 
```shell
docker run -p 8001:8001 --name grpc-hello chinlying/grpc-hello:V1.0
```

+ testing
```shell
cd client
go build && ./client
```

+ please refer to [common docker commands](../docker/commands.md) for more docker operations