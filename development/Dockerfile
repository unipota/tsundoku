FROM golang:1.13.4-stretch as golang
WORKDIR /tsundoku/server
RUN apt-get install -q git \
  &&  go get -u github.com/pilu/fresh
COPY go.mod ./
COPY go.sum ./
RUN go mod download

WORKDIR /tsundoku/server/ogp_canvas
COPY ogp_canvas/package*.json ./
RUN apt-get update -q \
    && apt-get install -qq build-essential libcairo2-dev libpango1.0-dev libjpeg-dev libgif-dev librsvg2-dev fontconfig
RUN curl -SL https://deb.nodesource.com/setup_13.x | bash
RUN apt-get install -y nodejs
RUN npm install --build-from-source
COPY ogp_canvas/font ./
COPY ogp_canvas/main.js ./

WORKDIR /tsundoku/server
ENTRYPOINT fresh -c fresh.conf