FROM golang:alpine

WORKDIR /build
COPY . .
RUN mv config.example.yaml config.yaml
RUN go build -o main main.go

CMD ["./main"]
