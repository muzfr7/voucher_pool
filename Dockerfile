FROM golang:1.16.5-buster

RUN apt-get update \
    && apt-get upgrade -y \
    git bash make gcc

RUN go get -u github.com/pressly/goose/cmd/goose

WORKDIR /voucher_pool

COPY ./ .
RUN go mod download
RUN make build

EXPOSE 8080

CMD ["./bin/main"]
