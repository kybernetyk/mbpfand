include $(GOROOT)/src/Make.inc

TARG=mbpfand
GOFILES=\
				config.go\
				sensor.go\
				logic.go\
				main.go

include $(GOROOT)/src/Make.cmd
