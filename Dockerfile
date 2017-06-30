FROM hub.c.163.com/library/golang:latest
MAINTAINER kaesalai@gmail.com
WORKDIR /gogogo/src/huanyu0w0
ENV GOPATH /gogogo
COPY . /gogogo/src/huanyu0w0
RUN go get -u github.com/labstack/echo
RUN go get -u github.com/dgrijalva/jwt-go
RUN go get -u github.com/russross/blackfriday
RUN go get -u qiniupkg.com/api.v7
RUN go get -u gopkg.in/mgo.v2
RUN go install huanyu0w0
ENTRYPOINT ["/gogogo/bin/huanyu0w0"]