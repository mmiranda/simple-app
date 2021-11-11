FROM golang:1.17.2 AS builder
WORKDIR /go/src/github.com/mmiranda/simple-app
RUN go get -d -v golang.org/x/net/html
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgop .

FROM alpine:latest
LABEL org.opencontainers.image.source https://github.com/mmiranda/simple-app
LABEL org.opencontainers.image.description "Simple app to use as example"
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /go/src/github.com/mmiranda/simple-app/simple-app .

VOLUME /data

ENTRYPOINT ["/app/simple-app"]
