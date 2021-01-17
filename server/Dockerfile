FROM golang:1.15.6 as builder

COPY . /app

WORKDIR /app

RUN go build -o server

FROM ubuntu:20.04

WORKDIR /app

COPY --from=builder /app/server .

ENV HOST=
ENV API_KEY=
ENV USER=
ENV PASS=

CMD ["./server"]