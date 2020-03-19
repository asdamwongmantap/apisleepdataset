#FROM golang:1.12-alpine
FROM golang:1.12
#RUN apk add --no-cache git

# Set the Current Working Directory inside the container

WORKDIR /apisleepdataset/

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod go.sum ./
# COPY go.sum .

RUN go mod download

COPY . .

COPY main.go .

USER 0:0
# Build the Go app

#USER 1001:0
# This container exposes port 8080 to the outside world
EXPOSE 2073

# Run the binary program produced by `go install`
CMD ["./main"]

