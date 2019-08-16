FROM golang:1.12.6-alpine
WORKDIR /tsundoku/ogp
RUN apk add --update --no-cache git \
  &&  go get -u github.com/pilu/fresh

RUN  apk add --update chromium

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ENTRYPOINT fresh -c fresh.conf