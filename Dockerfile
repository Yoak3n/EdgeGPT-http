FROM golang:alpine

WORKDIR /app
COPY . .
RUN mv config.example.yml config.yml
RUN touch cookies.json
RUN go build -o main main.go

CMD ["./main"]
