GONUMBER=6
COMPILER=$(GONUMBER)g
LINKER=$(GONUMBER)l

include $(GOROOT)/src/Make.inc

TARG=main
GOFILES=\
		10.go\
		37.go

run:
	$(COMPILER) 37.go 10.go
	$(LINKER) 37.$(GONUMBER)
	./$(GONUMBER).out
	@make clean

include $(GOROOT)/src/Make.pkg
