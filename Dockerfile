FROM python:3.10-slim as python_build

RUN apt-get update && apt-get install -y \
    build-essential && \
    pip install --no-cache-dir scrapyd-client

COPY /general_spider /src
WORKDIR /src
RUN scrapyd-deploy --build-egg general.egg

FROM node:current-alpine3.20 AS node_builder

COPY /general_spider_dashboard /src

WORKDIR /src

RUN npm install -g tailwindcss
RUN npm install -g clean-css-cli
RUN npx tailwindcss -i ./public/input.css -o ./tmp/output.css
RUN cleancss -o ./public/output.css ./tmp/output.css

FROM golang:1.23.4-alpine3.20 AS go_builder

RUN apk update && apk upgrade && apk add --no-cache ca-certificates tzdata

COPY /general_spider_dashboard /src
COPY --from=node_builder /src/public /src/public
COPY --from=python_build /src/general.egg /src/app/general.egg

WORKDIR /src

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
