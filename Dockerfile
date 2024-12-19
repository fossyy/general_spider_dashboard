FROM python:3.10-slim as python_build

RUN apt-get update && apt-get install -y \
    git \
    openssh-client \
    libssl-dev \
    libssh2-1-dev \
    build-essential && \
    pip install --no-cache-dir scrapyd-client

RUN git clone https://github.com/fossyy/general_spider.git
WORKDIR /general_spider
RUN scrapyd-deploy --build-egg general.egg

FROM node:current-alpine3.20 AS node_builder

WORKDIR /src
COPY /public /src/public
COPY tailwind.config.js .
COPY /view /src/view

RUN npm install -g tailwindcss
RUN npm install -g clean-css-cli
RUN npx tailwindcss -i ./public/input.css -o ./tmp/output.css
RUN cleancss -o ./public/output.css ./tmp/output.css

FROM golang:1.23.4-alpine3.20 AS go_builder

WORKDIR /src
COPY . .
COPY --from=node_builder /src/public /src/public
COPY --from=python_build /general_spider/general.egg /src/app/general.egg

RUN apk update && apk upgrade && apk add --no-cache ca-certificates tzdata
RUN update-ca-certificates
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate
RUN go build -o ./tmp/main

FROM scratch

WORKDIR /general

COPY --from=go_builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=go_builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=go_builder /src/public /general/public
COPY --from=go_builder /src/tmp/main /general

ENV TZ Asia/Jakarta

ENTRYPOINT ["./main"]