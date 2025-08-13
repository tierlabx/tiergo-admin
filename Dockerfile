FROM golang:1.24.6-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest


COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .


EXPOSE 88
ENTRYPOINT ["air","-c",".air.toml"]
