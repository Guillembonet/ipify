# Build the service to a binary
FROM golang:1.19.5-alpine AS builder

# Install packages
RUN apk add --no-cache bash gcc musl-dev linux-headers git

# Compile application
WORKDIR /go/src/github.com/guillembonet/ipify
ADD . .
RUN go build -o build/main main.go

# Copy and run the made binary
FROM alpine:3.17

COPY --from=builder /go/src/github.com/guillembonet/ipify/build/main /usr/bin/ipify

ENTRYPOINT ["/usr/bin/ipify"]