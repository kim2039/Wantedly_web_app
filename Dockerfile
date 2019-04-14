FROM golang:1.8

COPY hello.go /go
WORKDIR /go

RUN go get github.com/ant0ine/go-json-rest/rest

EXPOSE 80

CMD ["go", "run", "hello.go"]
