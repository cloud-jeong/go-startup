VER=1.0.0
GOARCH=amd64

all: clean darwin linux

darwin:
	echo "Make darwin binary ..."
	GOOS=darwin GOARCH=${GOARCH} go build -o ~/gows/bin/darwin/simpleweb_darwin_${VER}
	ln -s ~/gows/bin/darwin/simpleweb_darwin_${VER} ~/gows/bin/darwin/simpleweb

linux:
	echo "Make linux binary ..."
	GOOS=linux GOARCH=${GOARCH} go build -o ~/gows/bin/linux/simpleweb_linux_${VER}
	ln -s ~/gows/bin/linux/simpleweb_linux_${VER} ~/gows/bin/linux/simpleweb


clean:
	rm -f ~/gows/bin/darwin/simpleweb_darwin_${VER}
	rm -f ~/gows/bin/darwin/simpleweb
	rm -f ~/gows/bin/linux/simpleweb_linux_${VER}
	rm -f ~/gows/bin/linux/simpleweb
