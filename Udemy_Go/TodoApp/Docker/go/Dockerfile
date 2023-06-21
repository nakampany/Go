FROM golang:1.19.1-alpine

WORKDIR /usr/local/go/src/app

RUN apk update && apk add git

RUN go get github.com/lib/pq

CMD ["go", "run", "main.go"]