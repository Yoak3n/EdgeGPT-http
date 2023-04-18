FROM golang:alpine

WORKDIR /build

COPY config.example.yaml config.yaml
RUN go build -o main main.go

CMD ["./main"]
