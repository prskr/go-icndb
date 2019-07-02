FROM golang:1.12.6-alpine as build

RUN apk add --no-cache git && \
    go get -u github.com/gobuffalo/packr/v2/packr2

WORKDIR /icndb

ADD ./ ./

RUN packr2 && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o icndb cmd/icndb-server/main.go

FROM alpine:3.10

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

COPY --from=build --chown=icndb:icndb /icndb/icndb /usr/local/bin/icndb

USER icndb
EXPOSE 8000

ENTRYPOINT ["/usr/local/bin/icndb"]
CMD ["--host=0.0.0.0", "--port=8000", "--scheme=http", "--log-level=info", "--log-format=json"]