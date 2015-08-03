.PHONY: daemon

all: build

install: binary systemd

build: daemon cli

cli:
	go build -o dist/emru github.com/zoli/emru/cmd/emru

daemon:
	go build -o dist/emrud github.com/zoli/emru/daemon

binary:
	install -m 755 -p dist/emrud /usr/bin/emrud
	install -m 755 -p dist/emru /usr/bin/emru

systemd:
	install -m 644 -p dist/systemd/emru.service /etc/systemd/system/emru.service
	install -m 644 -p dist/systemd/emru.socket /etc/systemd/system/emru.socket

cli_install:
	go install github.com/zoli/emru/cmd/emru

test:
	go test ./...

clean:
	rm -r dist/emrud dist/emru
