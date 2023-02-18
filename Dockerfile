FROM golang:1.19.5-alpine3.17 as builder

WORKDIR /app

ENV CGO_BUILD=0
ENV GOOS=linux
ENV GOARCH=arm64


COPY . .

RUN go build main.go

FROM alpine

COPY --from=builder app/main app/main

CMD ["./app/main"]
