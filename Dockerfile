FROM golang
maintainer kaesalai@gmail.com
WORKDIR /gogogo/src/huanyu0w0
ENV GOPATH /gogogo
COPY . /gogogo/src/huanyu0w0
RUN go get github.com/labstack/echo
RUN go get github.com/russross/blackfriday
RUN go get qiniupkg.com/api.v7
RUN go get gopkg.in/mgo.v2
RUN cd /gogogo/src
RUN go install huanyu0w0
entrypoint ["/gogogo/bin/huanyu0w0"]