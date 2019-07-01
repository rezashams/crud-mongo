FROM golang:1.12.0-alpine3.9
# We create an /app directory within our
# image that will hold our application source
# files
RUN apk add git
RUN mkdir /app
# We copy everything in the root directory
# into our /app directory
ADD . /app

# We specify that we now wish to execute
# any further commands inside our /app
# directory
WORKDIR /app
Run go get ./...
# we run go build to compile the binary
# executable of our Go program
RUN go build -o web/main/main web/main/main.go
# Our start command which kicks off
# our newly created binary executable
CMD ["./app/web/main/main"]