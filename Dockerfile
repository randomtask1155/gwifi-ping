FROM golang:1.12 AS builder
LABEL maintainer="Daniel Lynch <danplynch@gmail.com>"
RUN mkdir -p /go/src/github.com/randomtask1155/gwifi-ping
WORKDIR $GOPATH/src/github.com/randomtask1155/gwifi-ping
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

ENV PORT=8080
ADD . .

RUN GOOS=linux GOARCH=arm GOARM=7 go build -o gwifi-ping .

FROM alpine:latest as certs
RUN apk --update add ca-certificates


FROM scratch
COPY --from=builder /go/src/github.com/randomtask1155/gwifi-ping/gwifi-ping /go/bin/gwifi-ping
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT ["/go/bin/gwifi-ping"]
