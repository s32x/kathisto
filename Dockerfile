FROM ubuntu:latest
RUN apt-get update && apt-get install -yq chromium-browser
ADD bin/server /usr/local/bin/server
CMD server