## All the docker commands

+ refer to [Docker CLI](https://docs.docker.com/engine/reference/run/)

## Common docker commands

+ [Build an image from a Dockerfile](https://docs.docker.com/engine/reference/commandline/build/)
```shell
docker build -t chinlying/grpc-hello:V1.0 .
```

+ [list images](https://docs.docker.com/engine/reference/commandline/images/)
```shell
docker images
```

+ [list containers](https://docs.docker.com/engine/reference/commandline/ps/)
```shell
docker ps # show running containers
docker ps -a # Show all containers (default shows just running)
```

+ [Create a new container, put a image into the container and then start the container](https://docs.docker.com/engine/reference/run/)
```shell
docker run -p 8001:8001 --name grpc-hello chinlying/grpc-hello:V1.0
```

+ [Start one or more stopped containers](https://docs.docker.com/engine/reference/commandline/start/)
```shell
docker start grpc-hello
```

+ [Stop one or more running containers](https://docs.docker.com/engine/reference/commandline/stop/)
```shell
docker stop grpc-hello
```

+ [Remove one or more containers](https://docs.docker.com/engine/reference/commandline/rm/)
```shell
docker rm grpc-hello
```