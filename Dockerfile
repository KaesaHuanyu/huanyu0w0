FROM hub.c.163.com/library/golang:latest
MAINTAINER kaesalai@gmail.com
WORKDIR /gogogo/src/huanyu0w0
ENV GOPATH /gogogo
COPY . /gogogo/src/huanyu0w0
RUN go get github.com/labstack/echo
RUN go get github.com/russross/blackfriday
RUN go get qiniupkg.com/api.v7
RUN go get gopkg.in/mgo.v2
    go install huanyu0w0
ENTRYPOINT ["/gogogo/bin/huanyu0w0"]