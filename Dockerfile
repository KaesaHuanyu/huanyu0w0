FROM golang:1.8
MAINTAINER kaesalai@gmail.com
WORKDIR /gogogo/src/huanyu0w0
ENV GOPATH /gogogo
COPY . /gogogo/src/huanyu0w0
RUN go get -u github.com/tools/godep
    godep save
    go install huanyu0w0
ENTRYPOINT ["/gogogo/bin/huanyu0w0"]