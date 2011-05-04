include $(GOROOT)/src/Make.inc

TARG=mbpfand
GOFILES=\
				config.go\
				read.go\
				main.go

include $(GOROOT)/src/Make.cmd
