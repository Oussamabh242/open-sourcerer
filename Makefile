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

build:
	$(GOCMD) build -o main
