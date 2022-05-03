# select alpine as the basic running environment
# It's better to refer to 'Dockerfile-Recommand' for learning how to write a better dockerfile
FROM alpine

# assign working directory
WORKDIR /go/app

# copy our executible binary file to the working directory
COPY ./server/server .

# expose port
EXPOSE 50051

# execute command
CMD ["./server"]