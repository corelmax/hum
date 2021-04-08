FROM golang:alpine

RUN apk add --no-cache \
    curl \
    git \
    make

# install dep
RUN OUTPUT_BIN=/usr/local/go/bin/dep \
    && curl --connect-timeout 60 -sSLo $OUTPUT_BIN https://github.com/golang/dep/releases/download/v0.5.1/dep-linux-amd64 \
    && chmod 755 $OUTPUT_BIN
RUN wget -O - https://github.com/corelmax/hum/archive/0.0.4.tar.gz | tar xvzf -
RUN mkdir /bin-hum
RUN cd hum* \
    && export GOPATH=$GOPATH:$PWD \
    && make \
    && cp ./bin/* /bin/
ENTRYPOINT [ "hum" ]