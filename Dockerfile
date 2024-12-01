FROM node:current-alpine3.20 AS node_builder

WORKDIR /src
COPY /public /src/public
COPY tailwind.config.js .
COPY /view /src/view

RUN npm install -g tailwindcss
RUN npm install -g clean-css-cli
RUN npx tailwindcss -i ./public/input.css -o ./tmp/output.css
RUN cleancss -o ./public/output.css ./tmp/output.css

FROM golang:1.23.1-alpine3.20 AS go_builder

WORKDIR /src
COPY . .
COPY --from=node_builder /src/public /src/public

RUN apk update && apk upgrade && apk add --no-cache ca-certificates tzdata
RUN update-ca-certificates
RUN go install github.com/a-h/templ/cmd/templ@$(go list -m -f '{{ .Version }}' github.com/a-h/templ)
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