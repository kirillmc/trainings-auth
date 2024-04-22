FROM golang:1.21.9-alpine3.18 AS builder

COPY . /github.com/kirillmc/trainings-auth/source/
WORKDIR /github.com/kirillmc/trainings-auth/source/

RUN go mod download
RUN go build -o ./bin/trainings_auth_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/kirillmc/trainings-auth/source/bin/trainings_auth_server .
COPY prod.env .
CMD ["./trainings_auth_server"]



#COPY .env /github.com/kirillmc/trainings-auth/source/bin/trainings_auth_server