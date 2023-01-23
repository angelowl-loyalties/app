FROM golang:alpine AS builder

RUN apk add --no-cache git

WORKDIR $GOPATH/src/

COPY . .

RUN go mod download
RUN go build -o /profiler

FROM alpine:3.17

RUN apk --no-cache add bash wget ca-certificates
RUN apk update

WORKDIR /

# do we need wait-for-it.sh?
COPY --from=builder /profiler .

ENTRYPOINT ["/profiler"]