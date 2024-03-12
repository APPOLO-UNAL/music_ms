# Use the official Golang image as a base
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Set environment variables
ENV SPOTIFY_CLIENT_ID=4bfedb806e3d49db9eeecfe03fe22057
ENV SPOTIFY_CLIENT_SECRET=a7c80d1d7d664c058c876c6290fc04db

# Copy all the files from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main ./app/cmd

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]