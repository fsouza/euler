GONUMBER=6
COMPILER=$(GONUMBER)g
LINKER=$(GONUMBER)l

include $(GOROOT)/src/Make.inc

TARG=main
GOFILES=\
		10.go\
		27.go

run:
	$(COMPILER) 27.go 10.go
	$(LINKER) 27.$(GONUMBER)
	./$(GONUMBER).out
	@make clean

include $(GOROOT)/src/Make.pkg
