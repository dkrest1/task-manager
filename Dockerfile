# Mini golang base image
FROM golang:1.20-alpine as builder

RUN apk add --no-cache build-base


# Create work dir
WORKDIR /app

# Copy go module files and download dependecy
COPY go.mod ./

COPY go.sum ./

RUN go mod download

# Copy the entire application to directory
COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o /app -a -ldflags '-linkmode external -extldflags "-static"' .

# Build the application
RUN go build -o main .

# Image for fina stage
FROM alpine:latest

# create work dir
WORKDIR /app

# Copy built app to work directory
COPY --from=builder /app/main .

# Copy env file
COPY .env .env

# Expose port
EXPOSE 9090

# Run the executable
CMD ["./main"]