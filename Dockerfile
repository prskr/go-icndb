FROM golang:1.11 as build

WORKDIR /go/src/github.com/baez90/go-icndb

ADD ./ ./

RUN go get -u github.com/gobuffalo/packr/packr && \
    packr && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o icndb .

FROM alpine:3.8

LABEL org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.name="ICNDB" \
      org.label-schema.description="Small and fast minimal ICNDB fork" \
      org.label-schema.url="https://github.com/baez90/go-icndb" \
      org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.vcs-url="https://github.com/baez90/go-icndb" \
      org.label-schema.vendor="" \
      org.label-schema.version="0.0.1" \
      org.label-schema.schema-version="1.0" \
      maintainer="peter.kurfer@gmail.com"

RUN adduser \
        -h /nonexistent \
        -g "" \
        -s /bin/false \
        -D \
        -H \
        -u 1000 \
        icndb

COPY --from=BUILD --chown=icndb:icndb /go/src/github.com/baez90/go-icndb/icndb /usr/local/bin/icndb

USER icndb
EXPOSE 8000

CMD ["/usr/local/bin/icndb"]