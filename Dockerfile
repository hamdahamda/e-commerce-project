# Gunakan base image untuk Golang
FROM golang:1.20

# Set working directory
WORKDIR /app

# Copy file ke dalam container
COPY . .

# Unduh dependencies
RUN go mod tidy

# Build aplikasi
RUN go build -o main .

# Jalankan aplikasi
CMD ["./main"]

# Expose port yang digunakan aplikasi (contoh: 8080)
EXPOSE 8080
