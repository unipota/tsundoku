FROM node:lts-alpine as client
WORKDIR /app
COPY client/package*.json ./
RUN npm install
COPY client .
RUN npm run build


FROM golang:1.12.6-alpine as server
WORKDIR /tsundoku/server

RUN apk add --update --no-cache git
COPY server/go.mod .
COPY server/go.sum .
RUN go mod download
COPY server .
RUN CGO_ENABLED=0 go build -o app


FROM alpine:3.9
WORKDIR /tsundoku
RUN apk add --update ca-certificates openssl && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*
ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz
EXPOSE 3000
COPY --from=client /app/dist ./static
COPY --from=server /tsundoku/server/app ./
COPY --from=server /tsundoku/server/static ./server/static

ENTRYPOINT ./app