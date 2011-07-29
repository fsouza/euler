GONUMBER=6
COMPILER=$(GONUMBER)g
LINKER=$(GONUMBER)l

include $(GOROOT)/src/Make.inc

TARG=main
GOFILES=\
		10.go\
		41.go

run:
	$(COMPILER) 41.go 10.go
	$(LINKER) 41.$(GONUMBER)
	./$(GONUMBER).out
	@make clean

include $(GOROOT)/src/Make.pkg
