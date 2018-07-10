FROM golang:1.9-alpine3.7 AS base

RUN apk -U upgrade && \
    apk --update add git && \
    go version

RUN cd /go/src && \
    go get -u github.com/golang/dep/cmd/dep

RUN cd /go/src && \
	git clone https://github.com/cloud-jeong/go-startup -b develop /go/src/go-startup && \
	rm -Rf go-startup/cmd/kcem && \
	rm -Rf go-startup/vendor && \
	rm -Rf go-startup/*GoPkg.*


RUN	cd /go/src/go-startup/cmd/simpleweb && \
	dep ensure && \
	go build -o simpleweb

FROM alpine:3.7

RUN mkdir -p /home/cocktail
RUN mkdir -p /home/cocktail
COPY --from=base /go/src/go-startup/cmd/simpleweb /home/cocktail

WORKDIR /home/cocktail
RUN mkdir log
ENV PORT 8000
EXPOSE $PORT

RUN chmod 755 /home/cocktail/simpleweb

CMD ["/home/cocktail/simpleweb"]
