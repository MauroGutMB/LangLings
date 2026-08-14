CC      ?= cc
CFLAGS  ?= -std=c17 -Wall -Wextra -O0 -g

test: testes.bin
	./testes.bin

testes.bin: exercicio.c testes.c exercicio.h
	$(CC) $(CFLAGS) -o $@ exercicio.c testes.c

.PHONY: test
