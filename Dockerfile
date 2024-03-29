
FROM golang:1.18-alpine

# RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV PORT=8080
ENV GO111MODULE=on

# Build the Go app
RUN go build -o . orderbook/orderbookapi/*.go

COPY orderbook/orderbookapi/swagger/swagger.yaml ./swagger/swagger.yaml


# This container exposes port 8080 to the outside world
EXPOSE $PORT

# Run the binary program produced by `go install`
CMD ["./main", "-env=prod"]