# Gunakan image Go yang sesuai
FROM golang:1.21 AS builder

WORKDIR /app

# Copy go.mod & go.sum dulu agar dependency cache tetap efisien
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh kode program
COPY . .

# Build aplikasi dengan CGO disabled agar lebih ringan
RUN CGO_ENABLED=0 go build -o app

# Stage kedua: gunakan image yang lebih ringan untuk runtime
FROM alpine:latest  

WORKDIR /root/

# Copy aplikasi yang sudah dibangun
COPY --from=builder /app/app .

# Jalankan aplikasi
CMD ["./app"]
