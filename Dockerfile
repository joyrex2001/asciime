
FROM docker.io/golang:1.10

ARG CODE=github.com/joyrex2001/asciime

ADD . /go/src/${CODE}/
ADD ./frontend /app/frontend
RUN cd /go/src/${CODE} && CGO_ENABLED=0 go build -o /app/main

WORKDIR /app
CMD ["./main"]
