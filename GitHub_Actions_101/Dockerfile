FROM golang:1.15-buster as build
WORKDIR /go/src/app
COPY . /go/src/app

ARG apiName

RUN go get -d -v ./...

RUN go build main.go

# FROM gcr.io/distroless/base-debian10
FROM golang:1.15-buster
COPY --from=build /go/src/app/main /
CMD ["/main"]