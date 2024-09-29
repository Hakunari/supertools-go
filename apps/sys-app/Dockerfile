FROM golang:latest

WORKDIR /app

# 复制依赖文件并下载模块依赖，分开复制依赖文件（go.mod 和 go.sum）和项目文件是为了优化构建过程。如果你在代码中做了任何更改，Docker 会重新构建镜像，但如果你分开复制依赖文件和项目文件，Docker 会利用缓存，加快构建速度。
COPY go.mod go.sum ./
RUN go mod download

# 复制项目文件到工作目录中
COPY . .

RUN go build -o sys-app ./cmd/main.go

CMD ["./sys-app"]
