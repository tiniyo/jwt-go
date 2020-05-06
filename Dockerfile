FROM golang:1.14 as builder
WORKDIR /go/src/github.com/jwt-go
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=builder /go/src/github.com/jwt-go/app .
COPY --from=builder /go/src/github.com/jwt-go/config /config

EXPOSE 8080

CMD ["/app"]
