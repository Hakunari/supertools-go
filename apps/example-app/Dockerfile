FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# 复制服务配置文件和启动脚本
COPY config.yaml /etc/config.yaml
COPY scripts/setup.sh /app/setup.sh

RUN go build -o sys-app main.go

CMD ["./sys-app"]
