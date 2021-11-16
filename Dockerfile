FROM golang:1.15 AS builder

WORKDIR /code/

COPY ./go.mod /code/go.mod
COPY ./go.sum /code/go.sum
RUN go mod download

COPY . /code/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./


FROM debian:stretch

COPY --from=builder /code/39_go_CI /usr/local/bin/39_go_CI

RUN chmod +x /usr/local/bin/39_go_CI

ENTRYPOINT [ "39_go_CI" ]