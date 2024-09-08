GOCMD=go
TEMPL=templ
BUILD_DIR=./tmp


generate:
	$(TEMPL) generate

run: generate
	$(GOCMD) run main.go

format: 
	$(GOCMD) fmt
	$(TEMPL) fmt .

build: generate
	$(GOCMD) build -tags netgo -ldflags '-s -w' -o app

