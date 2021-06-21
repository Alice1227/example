FROM golang:1.15 AS builder
RUN mkdir -p /app
WORKDIR /app
COPY . .

FROM alpine:3.12
WORKDIR /app
COPY --from=builder /app/example .
COPY --from=builder /app/env ./env
EXPOSE 80
ENTRYPOINT ["./example"]