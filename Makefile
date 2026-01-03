#
.PHONY:	dummy

GOCMD := $(shell command -v go)


tidy:	dummy
	$(RM) *~



build:
	$(GOCMD) build
