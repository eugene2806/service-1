FROM golang:1.23.3-alpine AS builder

COPY . /service-1/

WORKDIR /service-1/

RUN go mod download

RUN go build -o ./bin/rest_server ./main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /service-1/bin/rest_server .

COPY --from=builder /service-1/.env .

COPY --from=builder /service-1/migrate ./migrate

ENTRYPOINT ["/bin/sh", "-c", "./rest_server migrate up & ./rest_server rest"]