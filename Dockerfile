FROM ubuntu:14.04
RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y software-properties-common && \
    apt-add-repository multiverse && \
    add-apt-repository ppa:no1wantdthisname/ppa && \
    apt-get update && \
    echo ttf-mscorefonts-installer msttcorefonts/accepted-mscorefonts-eula select true | sudo debconf-set-selections && \
    apt-get install -y wget libfontconfig libfreetype6 libfreetype6-dev ttf-mscorefonts-installer fontconfig-infinality
ADD phantomjs /usr/local/bin/phantomjs
ADD render.js /render.js
ADD kathisto /usr/local/bin/kathisto
CMD kathisto
