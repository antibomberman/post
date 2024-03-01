FROM golang:1.21
WORKDIR /app
COPY . .

RUN go mod tidy

RUN go build -o app cmd/app/main.go
CMD ["./app"]