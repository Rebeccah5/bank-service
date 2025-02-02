# Use the official Golang image as a build stage
FROM golang:1.20 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o bank-service cmd/main.go

# Use a lightweight image for final container
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/bank-service .

EXPOSE 8080
CMD ["./bank-service"]
