FROM alpine:latest
WORKDIR /app
EXPOSE 80

#https://stackoverflow.com/questions/52056387/how-to-install-go-in-alpine-linux
ARG GOLANG_VERSION=1.21.7

RUN apk update && apk add go gcc bash musl-dev openssl-dev ca-certificates && update-ca-certificates

RUN wget https://dl.google.com/go/go$GOLANG_VERSION.src.tar.gz && tar -C /usr/local -xzf go$GOLANG_VERSION.src.tar.gz

RUN cd /usr/local/go/src && ./make.bash

ENV PATH=$PATH:/usr/local/go/bin

RUN rm go$GOLANG_VERSION.src.tar.gz

RUN apk del go

RUN go version

COPY main.go ./
ENV GO111MODULE=off
RUN go build -o rinha-app .

CMD [ "./rinha-app" ]