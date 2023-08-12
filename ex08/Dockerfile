FROM golang:1.20-alpine as builder
RUN mkdir /build
WORKDIR /build
COPY . .
ENV GOOS=linux CGO_ENABLED=0
RUN set -ex && \
    apk add --no-progress --no-cache \
    gcc \
    musl-dev

RUN go build -o server cmd/main.go

FROM alpine:3.15.0
WORKDIR /app
COPY --from=builder /build/server .
COPY .env .

CMD [ "/app/server" ]
