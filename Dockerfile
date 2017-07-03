FROM hub.c.163.com/library/golang:latest
MAINTAINER kaesalai@gmail.com
WORKDIR /gogogo/src/huanyu0w0
ENV GOPATH /gogogo
ENV TZ Asia/Shanghai
COPY . /gogogo/src/huanyu0w0
RUN go get -u github.com/labstack/echo && \
    go get -u github.com/dgrijalva/jwt-go && \
    go get -u github.com/russross/blackfriday && \
    go get -u qiniupkg.com/api.v7 && \
    go get -u gopkg.in/mgo.v2 && \
    go install huanyu0w0
ENTRYPOINT ["/gogogo/bin/huanyu0w0"]