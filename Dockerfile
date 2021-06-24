FROM golang:latest
RUN go env -w GO111MODULE=off
RUN go get github.com/gin-gonic/gin
RUN go get gopkg.in/olahol/melody.v1
RUN echo $GOROOT && echo $GOPATH
RUN git clone https://github.com/munvei/ase2021-group3.git /go/web_app
RUN cd /go/web_app && git checkout develop
ENTRYPOINT go run /go/web_app/ws_ren/main.go
