GOCMD=go
TEMPL=templ
BUILD_DIR=./tmp


generate:
	$(TEMPL) generate

run: generate
	$(GOCMD) run main.go


	
