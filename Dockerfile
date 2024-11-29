# Gunakan base image golang
FROM golang:1.19

# Setel direktori kerja dalam container
WORKDIR /app

# Copy go.mod dan go.sum files
COPY go.mod go.sum ./

# Unduh dependensi
RUN go mod download

# Copy seluruh kode sumber
COPY . .

# Build aplikasi Go
RUN go build -o myapp

# Expose port yang digunakan oleh aplikasi
EXPOSE 8080

# Jalankan aplikasi
CMD ["./myapp"]
