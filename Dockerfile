FROM ubuntu:latest
RUN apt-get update && apt-get install -yq chromium-browser
ADD kathisto /usr/local/bin/kathisto
CMD kathisto