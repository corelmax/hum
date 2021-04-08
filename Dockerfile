FROM golang:1.16-alpine

RUN apk add --no-cache \
    curl \
    git \
    make

# RUN mkdir -p /hum
# RUN curl -s -L https://github.com/corelmax/hum/archive/0.0.5.tar.gz | tar -xvz -C /hum

# RUN cp /hum/hum-0.0.5/bin/* /bin/
WORKDIR /hum
ADD . .
RUN make
RUN cp ./bin/* /bin/
ENTRYPOINT [ "sh", "-c" ]