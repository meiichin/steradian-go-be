FROM golang:1.23.2
LABEL maintainer="Martinus Yudi Purwono"

WORKDIR /app

# Copy go mod and sum files version and license
COPY go.mod go.sum ./
COPY version.json ./
COPY LICENSE ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Command to run
CMD ["./main"]