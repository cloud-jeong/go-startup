VER=1.0.0
GOARCH=amd64

all: clean darwin linux

darwin:
	echo "Make darwin binary ..."
	GOOS=darwin GOARCH=${GOARCH} go build -o ~/github/bin/darwin/simpleweb_darwin_${VER}
	ln -s ~/github/bin/darwin/simpleweb_darwin_${VER} ~/github/bin/darwin/simpleweb

linux:
	echo "Make linux binary ..."
	GOOS=linux GOARCH=${GOARCH} go build -o ~/github/bin/linux/simpleweb_linux_${VER}
	ln -s ~/github/bin/linux/simpleweb_linux_${VER} ~/github/bin/linux/simpleweb

clean:
	rm -f ~/github/bin/darwin/simpleweb_darwin_${VER}
	rm -f ~/github/bin/darwin/simpleweb
	rm -f ~/github/bin/linux/simpleweb_linux_${VER}
	rm -f ~/github/bin/linux/simpleweb
