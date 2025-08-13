FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/cosmtrek/air@latest

EXPOSE 88
ENTRYPOINT ["air","-c",".air.toml"]
