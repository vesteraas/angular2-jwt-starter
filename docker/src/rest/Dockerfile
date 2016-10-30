FROM golang:1.7.3

ENV GOPATH /go

RUN mkdir -p /go/src/rest
WORKDIR /go/src/rest

COPY . /go/src/rest

RUN go get
RUN go install

EXPOSE 8080

ENTRYPOINT /go/bin/rest
