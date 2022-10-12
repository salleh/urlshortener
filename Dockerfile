FROM golang:alpine AS builder

ARG LISTEN_PORT
ENV LISTEN_PORT=${LISTEN_PORT}

WORKDIR "/app"

# Copy "." (all folder and file in current directory) to "." (Docker image)
COPY . .
COPY .env .env

# Get the packages; "-d" is for stop the command from doing "go install"
RUN go get -d -v ./... 

# Install the packages
# RUN go install -v ./...

# Build the executable using make (from Makefile)
RUN go build -o url-shortener cmd/server/url-shortener.go

# Set the server's OS
FROM alpine:latest

# Copy the binary and control script from the builder stage to the live image
COPY --from=builder /app/url-shortener .
RUN ["chmod", "+x", "url-shortener"]

# Copy configurations and asset files
COPY --from=builder /app/.env .

# Run the application
# CMD ["./start-otp-server.sh"]
CMD ["./url-shortener"]

# Expose the container port
EXPOSE $LISTEN_PORT