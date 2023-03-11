FROM golang:alpine

WORKDIR /app

RUN apk add --no-cache git

RUN go install github.com/githubnemo/CompileDaemon@v1.4.0

COPY . .

RUN chmod +x ./build_app.sh

CMD ["CompileDaemon", "-build", "./build_app.sh", "-command", "./service_executable", "-polling=true"]