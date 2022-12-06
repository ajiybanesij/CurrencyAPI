FROM golang:1.18-alpine as builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o service .

FROM alpine:3.11.3
WORKDIR /usr/src/app
COPY --from=builder /build/service .
COPY . /usr/src/app

EXPOSE 8081
ENTRYPOINT [ "./service" ]
