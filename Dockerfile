FROM golang:1.20-alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 设置工作目录
WORKDIR /build

# 复制 Go 模块文件并下载依赖项
COPY go.mod go.sum ./
RUN go mod download

# 复制所有源码到工作目录
COPY . .

# 构建 Go 可执行文件
RUN go build -o main .

# 第二阶段：运行阶段
FROM alpine:latest

WORKDIR /root

# 复制配置文件夹
COPY  /config ./config

# 把静态文件拷贝到当前目录
COPY /template ./template

# 从拷贝到当前目录
COPY /wait-for-it.sh .

# 从编译阶段复制构建好的可执行文件
COPY --from=builder /build/main .

RUN set -eux; \
    apk update; \
    apk add --no-cache bash netcat-openbsd; \
    apk add --no-cache dos2unix && dos2unix wait-for-it.sh;\
    chmod 755 wait-for-it.sh



# 暴露应用程序的端口
#EXPOSE 8080

# 启动 Go 应用程序
#CMD ["./main"]