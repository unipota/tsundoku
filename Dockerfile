FROM node:13.1-alpine as client
WORKDIR /app
COPY client/package*.json ./
RUN npm install
COPY client .
RUN npm run build


FROM golang:1.13.4-alpine as server
WORKDIR /tsundoku/server
RUN apk add --update --no-cache git
COPY server/go.mod .
COPY server/go.sum .
RUN go mod download
COPY server/model ./model
COPY server/router ./router
COPY server/static ./static
COPY server/main.go ./
RUN CGO_ENABLED=0 go build -o app


FROM debian:stretch-slim

WORKDIR /tsundoku/ogp_canvas
COPY server/ogp_canvas/package*.json ./
RUN apt-get update -q \
    && apt-get install -qq build-essential libcairo2-dev libpango1.0-dev libjpeg-dev libgif-dev librsvg2-dev fontconfig
RUN curl -SL https://deb.nodesource.com/setup_13.x | bash
RUN apt-get install -y nodejs
RUN npm install --build-from-source
COPY server/ogp_canvas/font ./
COPY server/ogp_canvas/main.js ./

WORKDIR /tsundoku
RUN apt-get update -q && apt-get install -y wget
ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz
EXPOSE 3000
COPY --from=client /app/dist ./static
COPY --from=server /tsundoku/server/app ./
COPY --from=server /tsundoku/server/static ./server/static

ENTRYPOINT ./app