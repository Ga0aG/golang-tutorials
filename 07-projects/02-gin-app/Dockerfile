FROM golang:1.21-alpine AS builder

WORKDIR /app

# 安装构建依赖
RUN apk add --no-cache git

# 复制依赖文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/server/main.go

# 最终镜像
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 从builder复制编译后的二进制文件
COPY --from=builder /app/server .

# 暴露端口
EXPOSE 8080

# 设置环境变量
ENV PORT=8080 \
    ENV=production

CMD ["./server"]
