# File: Dockerfile
# File Created: 06 Mar 2019 11:31
# By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

FROM golang:onbuild
ADD . /go/src/server
WORKDIR /go/src/server
RUN go get server
RUN go install
ENTRYPOINT ["/go/bin/server"]