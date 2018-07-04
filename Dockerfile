FROM golang:1.10 AS builder

WORKDIR /go/src/tdrug

COPY . .

ENV CGO_ENABLED=0

RUN rm -f tdrug && go build tdrug.go

FROM alpine

COPY --from=builder /go/src/tdrug/tdrug /usr/local/bin/tdrug
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip

RUN echo 'tdrug:x:1000:1000::/:/sbin/nologin' >> /etc/passwd && \
    echo 'tdrug:!::0:::::' >> /etc/shadow

USER tdrug

CMD [ "/usr/local/bin/tdrug" ]
