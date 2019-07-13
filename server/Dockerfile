FROM golang:1.12.6-alpine
WORKDIR /tsundoku/server

RUN apk add --update --no-cache git
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o app
ENTRYPOINT [ "./app" ]
