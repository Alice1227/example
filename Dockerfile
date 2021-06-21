FROM golang:1.15-alpine AS builder
RUN mkdir -p /app
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build

FROM scratch
WORKDIR /app
COPY --from=builder /app/example .
COPY --from=builder /app/env ./env
EXPOSE 80
ENTRYPOINT ["./example"]
