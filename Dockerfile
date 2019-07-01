FROM golang:1.12.0-alpine3.9
# We create an /app directory within our
# image that will hold our application source
# files
RUN apk add git
# Make the source code path
RUN mkdir -p /go/src/github.com/rezashams/repository

# Add all source code
ADD . /go/src/repository
WORKDIR /go/src/repository
RUN go get ./...
# we run go build to compile the binary
# executable of our Go program
RUN go build -o main web/main/main.go
# Our start command which kicks off
# our newly created binary executable
CMD ["./main"]