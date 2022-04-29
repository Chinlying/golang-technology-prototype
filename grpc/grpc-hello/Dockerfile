# select ubuntu 20.04 as the basic running environment
FROM ubuntu:20.04

# assign working directory
WORKDIR /go/app

# copy our executible binary file to the working directory
COPY ./server/server .

# expose port
EXPOSE 8001

# execute command
CMD ["./server"]