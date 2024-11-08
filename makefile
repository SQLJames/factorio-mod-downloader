# Determine the operating system
ifeq ($(OS),Windows_NT)
    # Use 'type' command in Windows
    READ_VERSION := type VERSION
	HOMEDIR := $(USERPROFILE)
else
    # Use 'cat' command in Unix-like systems
    READ_VERSION := cat VERSION
	HOMEDIR := $(HOME)
endif

VERSION := $(shell $(READ_VERSION))

debug:
	go run .\cmd\cli\main.go

build:
	docker run --rm -v ".:/work" -w /work/cmd/factorio-mod-downloader -e GOOS="$(go env GOOS)" -e GOARCH="$(go env GOARCH)" cgr.dev/chainguard/go build -o /work/bin/factorio-mod-downloader .
	docker run --rm -v ".:/work" -w /work/cmd/factorio-mod-downloader -e GOOS="windows" -e GOARCH="$(go env GOARCH)" cgr.dev/chainguard/go build -o /work/bin/factorio-mod-downloader.exe .
docker:
	@echo "Building docker image for factorio-mod-downloader version $(VERSION)"
	docker build -t factorio-mod-downloader:$(VERSION) .

docker-run:
	@echo "Running docker image for factorio-mod-downloader version $(VERSION)"
	docker run -it --rm factorio-mod-downloader:$(VERSION) download