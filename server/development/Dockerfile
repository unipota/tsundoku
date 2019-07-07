FROM golang:1.12.6-alpine
WORKDIR /tsundoku/server
RUN apk add --update --no-cache git \
  &&  go get -u github.com/pilu/fresh

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o tsundoku main.go

COPY development/fresh.conf ./
ENTRYPOINT fresh -c fresh.conf