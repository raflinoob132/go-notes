# Gunakan image Golang sebagai base image
FROM golang:1.21

# Set environment variable
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Buat direktori kerja dalam container
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Download dependensi dan build aplikasi
RUN go mod tidy && go build -o app

# Jalankan aplikasi saat container di-start
CMD ["/app/app"]
