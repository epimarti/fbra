FROM golang

ADD . /go/src/github.com/epimarti/fbra

RUN go get -u github.com/gorilla/mux

RUN go install github.com/epimarti/fbra

ENTRYPOINT /go/bin/fbra

EXPOSE 80
