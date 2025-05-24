.PHONY: dev build run clean

ifneq (,$(wildcard .env))
    include .env   # imports KEY=VALUE lines as Make variables :contentReference[oaicite:0]{index=0}
    export         # exports all Make variables into the environment for recipes :contentReference[oaicite:1]{index=1}
endif

dev:
	air

build:
	go build -o bin/server ./cmd/server

run:
	./bin/server

clean:
	rm -rf bin tmp
