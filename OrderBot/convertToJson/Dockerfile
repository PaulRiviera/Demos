FROM golang:1.15-buster as build
WORKDIR /go/src/app
COPY . /go/src/app
RUN go build ./main.go
EXPOSE ${PORT}
CMD ["/go/src/app/main"]
