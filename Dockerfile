FROM golang:1.20-alpine as builder

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/app/kubetelebot ./cmd/main.go

FROM alpine:3

LABEL org.opencontainers.image.authors="ashwinath@hotmail.com"
LABEL org.opencontainers.image.source https://github.com/ashwinath/kubetelebot

WORKDIR /usr/src/app
COPY --from=builder /usr/local/app/kubetelebot ./kubetelebot
COPY --from=bitnami/kubectl:1.27.7 /opt/bitnami/kubectl/bin/kubectl /usr/bin/kubectl

CMD ["./kubetelebot"]
