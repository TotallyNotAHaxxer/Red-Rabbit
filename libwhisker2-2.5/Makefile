# this is a passthru makefile for libwhisker2

DESTDIR=

lib:
	perl Makefile.pl lib

build:
	perl Makefile.pl lib

install:
	export DESTDIR
	perl Makefile.pl install

clean:
	perl Makefile.pl clean

nopod:
	perl Makefile.pl nopod

uninstall:
	export DESTDIR
	perl Makefile.pl uninstall
