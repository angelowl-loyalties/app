FROM golang:1.19-alpine3.17

RUN apk add --no-cache git

RUN go install github.com/githubnemo/CompileDaemon@v1.4.0

WORKDIR /app

COPY . .

RUN chmod +x build_app.sh

CMD ["CompileDaemon", "-build", "./build_app.sh", "-command", "./service_executable", "-polling=true"]