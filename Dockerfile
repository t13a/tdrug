FROM golang:1.10 AS builder

WORKDIR /go/src/tdurl

COPY . .

ENV CGO_ENABLED=0

RUN rm -f tdurl && go build tdurl.go

FROM alpine

COPY --from=builder /go/src/tdurl/tdurl /usr/local/bin/tdurl
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip

RUN echo 'tdurl:x:1000:1000::/:/sbin/nologin' >> /etc/passwd && \
    echo 'tdurl:!::0:::::' >> /etc/shadow

USER tdurl

CMD [ "/usr/local/bin/tdurl" ]
