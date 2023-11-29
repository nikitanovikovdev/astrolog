FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY . .

COPY .env.docker .env

COPY ./migrations /migrations

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/.env .
COPY --from=builder /app/main .
COPY --from=builder /app/migrations migrations

EXPOSE 5632

CMD ["./main"]
