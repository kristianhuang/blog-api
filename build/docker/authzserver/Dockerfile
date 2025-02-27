FROM golang:1.17.7-alpine3.15 AS builder

ENV GOPROXY=https://goproxy.cn \
    CGO_ENABLED=0 \
    GOOS=linux \
    TZ=Asia/Shanghai

WORKDIR /app

COPY go.mod .

RUN go mod tidy

COPY . .

RUN go build -ldflags="-w -s" -a -installsuffix cgo -o blog-authzserver ./cmd/authzserver

FROM scratch

ENV APISERVERPORT=8080 \
    TZ=Asia/Shanghai

WORKDIR /blog

# 从 builder 镜像中二进制文件拷贝到目录中
COPY --from=builder /app/blog-authzserver /blog/

COPY --from=builder /app/config/blog-authzserver.yaml /blog/

EXPOSE $APISERVERPORT

# 需要运行的命令
ENTRYPOINT ["/blog/blog-authzserver"]
CMD ["-c", "/blog/blog-authzserver.yaml"]