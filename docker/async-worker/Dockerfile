FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build-async-worker
RUN chmod -R o+rx /app/tmp/build

FROM debian:stable-slim

WORKDIR /app
COPY --from=builder /app/tmp/build /app/bin

CMD ["/app/bin/async_worker"]
