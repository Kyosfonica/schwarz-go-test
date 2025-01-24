# build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

WORKDIR /app/cmd/coupon_service
RUN go build -o main.go

ENTRYPOINT ["/app/cmd/coupon_service/main"]
EXPOSE 8080