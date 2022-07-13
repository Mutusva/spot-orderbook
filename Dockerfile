FROM golang:1.12-alpine

RUN apk add --no-cache git

C

# Set the Current Working Directory inside the container
WORKDIR /app/orderbookapi

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY ./src/orderbookapi/cmd .

# Build the Go app
RUN go build -o ./app/orderbookapi .


# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./app/orderbookapi"]