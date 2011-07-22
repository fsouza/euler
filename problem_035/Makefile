GONUMBER=6
COMPILER=$(GONUMBER)g
LINKER=$(GONUMBER)l

include $(GOROOT)/src/Make.inc

TARG=main
GOFILES=\
		10.go\
		35.go

run:
	$(COMPILER) 35.go 10.go
	$(LINKER) 35.$(GONUMBER)
	./$(GONUMBER).out
	@make clean

include $(GOROOT)/src/Make.pkg
