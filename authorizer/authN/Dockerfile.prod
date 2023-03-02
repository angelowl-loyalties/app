FROM golang:1.19-alpine3.17 AS builder

WORKDIR /app

COPY . .

RUN go build -o /service_executable

FROM alpine:3.17

COPY --from=builder /service_executable /service_executable

ENTRYPOINT ["/service_executable"]