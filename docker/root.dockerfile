FROM golang:1.24-alpine3.22 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY cmd/root-service/main.go .

RUN go build -o root-service .

FROM alpine:3.22

RUN apk --no-cache add ca-certificates 

WORKDIR /app

COPY --from=builder /app/root-service .

CMD [ "./root-service" ]