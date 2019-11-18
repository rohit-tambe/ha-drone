# First stage
FROM golang:1.13

# Make app directory
RUN mkdir /app

# Add app directory 
ADD . /app

# Create and change to the app directory.
WORKDIR /app

# Copy local code to the container image.
ADD . /go/src/github.com/ha-drone/

# Add lib
RUN go get github.com/go-playground/validator
RUN go get github.com/labstack/echo

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Second stage
FROM alpine:latest

# Run the web service on container startup.
CMD ["/app"]

# Copy from first stage
COPY --from=0 /app/main /app