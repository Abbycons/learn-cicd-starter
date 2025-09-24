# Builder stage
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Build for amd64 (Railway servers)
RUN GOOS=linux GOARCH=amd64 go build -o notely .

# Final stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/notely .
EXPOSE 8080
CMD ["./notely"]
