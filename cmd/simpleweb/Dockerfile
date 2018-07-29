FROM golang:1.9-alpine3.7 AS base

RUN apk -U upgrade && \
    apk --update add git && \
	go get -u github.com/golang/dep/cmd/dep && \
	git clone https://github.com/cloud-jeong/go-startup -b develop /tmp/go-startup && \
	cp -R /tmp/go-startup/cmd/simpleweb /go/src && \
	cd /go/src/simpleweb && \
	rm -f Gopkg.* &&\
	rm -Rf vendor && \
	dep init && \
	dep ensure && \
	go build -o app

FROM alpine:3.7

RUN mkdir -p /home/cocktail
COPY --from=base /go/src/simpleweb /home/cocktail/simpleweb

WORKDIR /home/cocktail/simpleweb
RUN mkdir log
ENV PORT 8000
EXPOSE $PORT

RUN chmod 755 /home/cocktail/simpleweb/app

CMD ["/home/cocktail/simpleweb/app"]
