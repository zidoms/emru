.PHONY: daemon gui

all: build

build: daemon cli gui

install: binary systemd desktop

cli:
	go build -o dist/emru github.com/zoli/emru/cmd/emru

daemon:
	go build -o dist/emrud github.com/zoli/emru/daemon

gui:
	cd gui && zip -r ../dist/gui/emru.nw ./* && cd ..
	cp /usr/lib/node-webkit/nw.pak dist/gui/nw.pak
	cp /usr/lib/node-webkit/icudtl.dat dist/gui/icudtl.dat
	cat /usr/lib/node-webkit/nw dist/gui/emru.nw > dist/gui/emrugui
	chmod +x dist/gui/emrugui
	rm dist/gui/emru.nw

binary:
	mkdir -p /usr/lib/emru
	install -m 755 -p dist/gui/emrugui /usr/lib/emru/emrugui
	install -m 644 -p dist/gui/nw.pak /usr/lib/emru/nw.pak
	install -m 644 -p dist/gui/icudtl.dat /usr/lib/emru/icudtl.dat
	install -m 755 -p dist/emrud /usr/lib/emru/emrud
	install -m 755 -p dist/emru /usr/lib/emru/emru
	ln -fs /usr/lib/emru/emru /usr/bin/emru

systemd:
	install -m 644 -p dist/systemd/emru.service /usr/lib/systemd/system/emru.service

cli_install:
	go install github.com/zoli/emru/cmd/emru

desktop:
	install -m 644 -p dist/gui/emru.desktop /usr/share/applications/emru.desktop

test:
	go test ./...

clean:
	rm -rf dist/emrud dist/emru ${GOPATH}/bin/emru
	rm -rf dist/gui/emrugui dist/gui/nw.pak dist/gui/icudtl.dat

uninstall:
	rm -f /usr/bin/emru /usr/lib/systemd/system/emru.service
	rm -rf /usr/lib/emru
