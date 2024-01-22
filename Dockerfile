FROM golang:1.21.6-alpine3.19 AS builder
LABEL authors="glodi"

#RUN apk --no-cache add ca-certificates
# WORKDIR /root/

# CMD ["./app"]

WORKDIR /go-proj

COPY go.mod go.sum ./

RUN go mod download

# FROM alpine:latest

COPY cmd/ ./
COPY configs/ ./
COPY internal/ ./
COPY pkg/ ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go-saas

RUN chmod +x /go-saas

ENTRYPOINT ["/go-saas"]
