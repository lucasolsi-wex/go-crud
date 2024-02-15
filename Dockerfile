FROM golang:1.21.5

WORKDIR /app

COPY . .

WORKDIR /app/cmd/user-api

RUN go mod download

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
