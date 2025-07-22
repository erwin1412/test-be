# Gunakan base image Golang resmi
FROM golang:1.24-alpine

# Install git & build tools (kalau perlu)
RUN apk update && apk add --no-cache git

# Set Working Directory
WORKDIR /app

# Copy go.mod & go.sum dulu (untuk cache layer)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy semua source code
COPY . .

# Build binary
RUN go build -o main .

# Expose port
EXPOSE 8082

# Jalankan binary
CMD ["./main"]
