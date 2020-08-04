# modget
# See LICENSE for copyright and license details.
.POSIX:

include config.mk

all: clean build

build:
	go build
	scdoc < modget.1.scd | sed "s/VERSION/$(VERSION)/g" > modget.1

clean:
	rm -f modget
	rm -f modget.1

install: build
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp -f modget $(DESTDIR)$(PREFIX)/bin
	chmod 755 $(DESTDIR)$(PREFIX)/bin/modget
	mkdir -p $(DESTDIR)$(MANPREFIX)/man1
	cp -f modget.1 $(DESTDIR)$(MANPREFIX)/man1/modget.1
	chmod 644 $(DESTDIR)$(MANPREFIX)/man1/modget.1

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/modget
	rm -f $(DESTDIR)$(MANPREFIX)/man1/modget.1

.PHONY: all build clean install uninstall
