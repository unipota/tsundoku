FROM nginx
COPY nginx.conf /etc/nginx/nginx.conf

FROM golang:1.12.6-alpine AS build
WORKDIR /tsundoku/server

RUN apk add --update --no-cache git
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /tsundoku/server
