FROM alpine:latest
RUN apk add --no-cache ca-certificates chromium
ADD kathisto /usr/local/bin/kathisto
CMD kathisto