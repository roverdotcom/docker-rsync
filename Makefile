GOOS="darwin"
# GOOS="linux"
# GOOS="windows"
GOARCH="amd64"

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) godep go build -o build/docker-rsync
	cp build/docker-rsync build/docker-rsync.v`./build/docker-rsync --version`.$(GOOS).$(GOARCH)
	(cd build && tar -cvzf docker-rsync.v`./docker-rsync -version`.$(GOOS).$(GOARCH).tar.gz docker-rsync.v`./docker-rsync -version`.$(GOOS).$(GOARCH))

clean:
	rm -r build/*

install:
	godep go install

test:
	GOTEST=1 go test -v dockermachine*


.PHONY: build clean test install
