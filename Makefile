build:
	test -f minisat.tar.gz || wget https://github.com/niklasso/minisat/tarball/master -O minisat.tar.gz
	tar xvf minisat.tar.gz
	cp -r niklasso-minisat-*/* minisat/
	rm -rf niklasso-minisat-*/
	rm minisat.tar.gz
	cd minisat/ && make
