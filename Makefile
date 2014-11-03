PWD = $(shell pwd)
NW  = $(PWD)/frontend/bin/emru/linux64
INSTALL_FILE    = install -m 644 -p
INSTALL_PROGRAM = install -m 755 -p
INSTALL_TARGET  = /usr/lib/emru/

all:

install_lib:
	mkdir -p $(INSTALL_TARGET)
	-$(INSTALL_FILE) $(NW)/nw.pak $(INSTALL_TARGET)
	-$(INSTALL_FILE) $(NW)/libffmpegsumo.so $(INSTALL_TARGET)
	-$(INSTALL_PROGRAM) $(NW)/emru $(INSTALL_TARGET)
install_bin:
	ln -s $(INSTALL_TARGET)/emru /usr/bin/emru
install_icons:
	-$(INSTALL_FILE) $(PWD)/frontend/app/icon/128/emru.png /usr/share/icons/hicolor/128x128/apps/
	-$(INSTALL_FILE) $(PWD)/frontend/app/icon/64/emru.png /usr/share/icons/hicolor/64x64/apps/
	-$(INSTALL_FILE) $(PWD)/frontend/app/icon/32/emru.png /usr/share/icons/hicolor/32x32/apps/
install_desktop:
	-$(INSTALL_FILE) $(PWD)/emru.desktop /usr/share/applications/

install: install_lib install_bin install_icons install_desktop
