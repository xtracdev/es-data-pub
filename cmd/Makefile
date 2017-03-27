image:
	docker run --rm -v $(CURDIR)/../:/go/src/github.com/xtracdev/es-data-pub -w /go/src/github.com/xtracdev/es-data-pub xtracdev/goora bash -c "cd cmd; make -f Makefile.docker"
	docker build -t xtracdev/pubevents .
