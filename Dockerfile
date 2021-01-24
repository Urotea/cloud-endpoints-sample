# Build Container
FROM golang:latest as builder
WORKDIR /src
COPY . .
# Set Environment Variable
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
# Build
WORKDIR /src/server
RUN go build -o app ./...

# Runtime Container
FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /src/server/app /app
ENTRYPOINT ["/app"]
