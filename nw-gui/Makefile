INSTALL_FILE = install -m 644 -p
INSTALL_PROGRAM = install -m 755 -p
INSTALL_TARGET = /usr/lib/emru
NODE_WEBKIT = bin/node-webkit-v0.10.5-linux-x64

all: frontend node_webkit build

frontend:
	cd app; npm install; ./node_modules/gulp/bin/gulp.js build;
	cd app; zip -r ../bin/app.nw package.json dist/*;

build:
	go build -o bin/emru
	cd bin; cat ./node-webkit-v0.10.5-linux-x64/nw app.nw > app && chmod +x app;
	mkdir -p $(HOME)/.local/share/emru

node_webkit:
	./install_nw.bash

install: install_lib install_bin install_icons install_desktop

install_lib:
	mkdir -p $(INSTALL_TARGET)
	-$(INSTALL_FILE) $(NODE_WEBKIT)/nw.pak $(INSTALL_TARGET)
	-$(INSTALL_FILE) $(NODE_WEBKIT)/nwsnapshot $(INSTALL_TARGET)
	-$(INSTALL_FILE) $(NODE_WEBKIT)/icudtl.dat $(INSTALL_TARGET)
	-$(INSTALL_PROGRAM) bin/app $(INSTALL_TARGET)
	-$(INSTALL_PROGRAM) bin/emru $(INSTALL_TARGET)

install_bin:
	ln -fs $(INSTALL_TARGET)/emru /usr/bin/emru

install_icons:
	-$(INSTALL_FILE) data/icons/128/emru.png /usr/share/icons/hicolor/128x128/apps/
	-$(INSTALL_FILE) data/icons/64/emru.png /usr/share/icons/hicolor/64x64/apps/
	-$(INSTALL_FILE) data/icons/32/emru.png /usr/share/icons/hicolor/32x32/apps/

install_desktop:
	-$(INSTALL_FILE) data/emru.desktop /usr/share/applications/

clean:
	cd bin; rm -rf emru app app.nw
	cd app/dist; rm -rf app.css app.js index.html
	rm -rf $(NODE_WEBKIT)