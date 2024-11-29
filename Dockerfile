# Gunakan base image untuk Golang
FROM golang:1.20-alpine

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Tentukan working directory
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Download module dependencies
RUN go mod download

# Build aplikasi
RUN go build -o main .

# Ekspos port yang akan digunakan
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]
