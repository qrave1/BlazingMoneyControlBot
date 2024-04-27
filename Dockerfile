FROM golang:1.22-alpine AS builder
LABEL authors="qrave1"

ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /moneyBot

COPY . .
RUN go mod download

RUN go build -o app cmd/main.go

FROM alpine as runtime

COPY --from=builder /moneyBot/.env /.env
COPY --from=builder /moneyBot/app /app
ENTRYPOINT["/app"]
