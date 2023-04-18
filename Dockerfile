FROM golang:alpine

WORKDIR /build
COPY go.mod go.mod
COPY go.sum go.sum
COPY config.example.yaml config.yaml
RUN go build -o main main.go

CMD ["./main"]
