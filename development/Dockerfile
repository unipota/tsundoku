FROM golang:1.13.4-stretch as golang

WORKDIR /tsundoku/server
COPY ogp_canvas/package*.json ogp_canvas/

WORKDIR /tsundoku/server/ogp_canvas
RUN apt-get update \
    && apt-get install -qq build-essential libcairo2-dev libpango1.0-dev libjpeg-dev libgif-dev librsvg2-dev fontconfig
RUN curl -SL https://deb.nodesource.com/setup_13.x | bash
RUN apt-get install -y nodejs
RUN npm install
COPY ogp_canvas/font ./
COPY ogp_canvas/main.js ./

WORKDIR /tsundoku/server
RUN apt-get install git \
  &&  go get -u github.com/pilu/fresh
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY model ./
COPY router ./
COPY static ./
COPY tmp ./
COPY fresh.conf ./
COPY main.go ./

ENTRYPOINT fresh -c fresh.conf