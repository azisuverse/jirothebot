# Gunakan image golang sebagai base
FROM golang:1.23-alpine

# Set working directory
WORKDIR /app

# Copy go.mod dan go.sum terlebih dahulu
COPY go.mod go.sum ./
RUN go mod tidy

# Copy seluruh kode source ke dalam container
COPY . .

# Build aplikasi Go
RUN go build -o bot ./cmd/bot

# Tentukan port yang digunakan
EXPOSE 8080

# Jalankan aplikasi
CMD ["./bot"]
