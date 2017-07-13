FROM golang:1.8.3-alpine3.6
MAINTAINER kaesalai@gmail.com
WORKDIR /gogogo/src/huanyu0w0
ENV GOPATH /gogogo
ENV TZ Asia/Shanghai
COPY . /gogogo/src/huanyu0w0
RUN go get -u github.com/labstack/echo-contrib/session && \
    go get -u github.com/gorilla/sessions && \
    go get -u github.com/afocus/captcha && \
    go get -u github.com/labstack/echo && \
    go get -u github.com/dgrijalva/jwt-go && \
    go get -u github.com/russross/blackfriday && \
    go get -u qiniupkg.com/api.v7 && \
    go get -u gopkg.in/mgo.v2 && \
    go install huanyu0w0
ENTRYPOINT ["/gogogo/bin/huanyu0w0"]