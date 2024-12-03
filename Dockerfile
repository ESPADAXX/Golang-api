# Use the official Go image as the base image
FROM golang:1.23.3

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Set Go proxy environment variable for reliable module fetching
ENV GOPROXY=https://proxy.golang.org,direct

# Download and cache the dependencies
RUN go mod download

# Copy the rest of the application code to the container
COPY . .

# Expose the port that the application will run on
EXPOSE 5050

# Set the default command to run the application
CMD ["go", "run", "cmd/server/main.go"]
