FROM golang:1.15.6 as builder

COPY . /app

WORKDIR /app

RUN go build -o configurator

FROM ubuntu:latest

WORKDIR /app

COPY --from=builder /app .

ENV HOST=
ENV API_KEY=
ENV USER=
ENV PASS=

CMD ["./configurator"]