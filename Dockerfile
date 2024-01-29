# Go image as the base image
FROM golang:1.20-alpine as builder

# Set working directory
WORKDIR /app

# Copy application code into the working directory
COPY . .

# Builid the application
RUN go build -o main

# Expose application port
EXPOSE ${PORT}

# Run executable
CMD [ "./main" ]
