FROM golang:1.15 as build-env

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./cmd/webui
RUN go build -v ./cmd/webui

FROM gcr.io/distroless/base-debian10

USER nobody:nobody
COPY --from=build-env /go/src/app/webui /
ADD ./html /html
EXPOSE 8000

CMD ["/webui"]