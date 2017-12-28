.PHONY: all windows darwin linux clean

version ?= dev
appname := robolearn
sources := build/web
build = docker run --rm -v $(PWD):/go/src/app:ro -w /go/src/app/server golang:1.8 bash -c "go get -u github.com/jteeuwen/go-bindata/...; go generate; go get; GOOS=$(1) GOARCH=$(2) go build -o ../build/$(appname)$(3)"
tar = cd build && tar -cvzf $(1)_$(2).tar.gz $(appname)$(3) && rm $(appname)$(3)
zip = cd build && zip $(1)_$(2).zip $(appname)$(3) && rm $(appname)$(3)

default: build/web build/robolearn


build_docker:
	docker build . -t robolearn:${version}

run:
	docker run -p 8080:80 robolearn:${version}

dev_node:
	cd web \
		&& npm install \
		&& npm start

dev_go:
	docker run -ti -v $$(pwd):/go/src/app -p 9000:9000 -w /go/src/app/server golang bash -c "go get; go get -u github.com/jteeuwen/go-bindata/...; go generate; go get github.com/tockins/realize; realize start"

dev_sdk:
	mkdir -p sdk
	mkdir -p sdk/python
	cd sdk/python; git clone git@github.com:NoUseFreak/python-robolearn.git . || git fetch

clean:
	rm -rf build/

build/robolearn:
	docker run --rm -v $(PWD):/go/src/app -w /go/src/app/server golang:1.8 bash -c "go get -u github.com/jteeuwen/go-bindata/...; go generate; go get; GOOS=darwin GOARCH=amd64 go build -o ../build/robolearn"

all: darwin linux windows


##### Common FRONTEND #####
build/web:
	docker run --rm -v "$(PWD)/web:/app" -w /app node bash -c "npm install; npm run build"
	mkdir -p build
	mv web/build build/web

##### DARWIN (MAC) BUILDS #####
darwin: build/darwin_amd64.tar.gz

build/darwin_amd64.tar.gz: $(sources)
	$(call build,darwin,amd64,)
	$(call tar,darwin,amd64)

##### LINUX BUILDS #####
linux: build/linux_arm.tar.gz build/linux_arm64.tar.gz build/linux_386.tar.gz build/linux_amd64.tar.gz

build/linux_386.tar.gz: $(sources)
	$(call build,linux,386,)
	$(call tar,linux,386)

build/linux_amd64.tar.gz: $(sources)
	$(call build,linux,amd64,)
	$(call tar,linux,amd64)

build/linux_arm.tar.gz: $(sources)
	$(call build,linux,arm,)
	$(call tar,linux,arm)

build/linux_arm64.tar.gz: $(sources)
	$(call build,linux,arm64,)
	$(call tar,linux,arm64)


##### WINDOWS BUILDS #####
windows: build/windows_386.zip build/windows_amd64.zip

build/windows_386.zip: $(sources)
	$(call build,windows,386,.exe)
	$(call zip,windows,386,.exe)

build/windows_amd64.zip: $(sources)
	$(call build,windows,amd64,.exe)
	$(call zip,windows,amd64,.exe)
