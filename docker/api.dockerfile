FROM golang:1.24-alpine3.22 AS builder

WORKDIR /app

COPY cmd/api-service/main.go .

RUN go build -o api-service /app/main.go

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/api-service .

CMD [ "/app/api-service" ]