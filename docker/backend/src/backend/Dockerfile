FROM golang:1.7.3

ENV GOPATH /go

RUN mkdir -p /go/src/backend
WORKDIR /go/src/backend

COPY . /go/src/backend

RUN go get
RUN go install

EXPOSE 8080

ENTRYPOINT /go/bin/backend
