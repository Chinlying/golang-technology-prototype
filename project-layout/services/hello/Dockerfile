FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /growing-into-an-excellent-golang-developer
ADD go.sum .
ADD go.mod .
COPY ./project-layout/services/hello ./project-layout/services/hello
RUN go build -o /app/server ./project-layout/services/hello/server/server.go


FROM alpine
WORKDIR /go/app
COPY --from=builder /app/server .
# expose port
EXPOSE 50051

CMD ["./server"]