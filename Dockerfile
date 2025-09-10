FROM golang:1.24.6-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest
# 安装 air
RUN apk add --no-cache git && \
    go install github.com/air-verse/air@latest

# 将 air 加入 PATH
ENV PATH=$PATH:/go/bin

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .


EXPOSE 88
ENTRYPOINT ["air","-c",".air.toml"]
