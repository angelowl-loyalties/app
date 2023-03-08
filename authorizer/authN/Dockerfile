FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o /service_executable authN.go

FROM public.ecr.aws/lambda/provided:al2

COPY --from=builder /service_executable /service_executable

ADD https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie /usr/bin/aws-lambda-rie
RUN chmod 755 /usr/bin/aws-lambda-rie

COPY authN/entry.sh /
RUN chmod 755 /entry.sh

ENTRYPOINT [ "/entry.sh" ]