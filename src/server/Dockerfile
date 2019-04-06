# File: Dockerfile
# File Created: 06 Mar 2019 11:31
# By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

FROM golang:1.8
ADD ./src/server /go/src/server
ADD ./public public
WORKDIR /go/src/server
RUN go build
ENTRYPOINT ["./server"]