GOCMD=go
TEMPL=templ
BUILD_DIR=./tmp
NPM=npm
NPX=npx

generate:
	$(TEMPL) generate

run: generate
	$(GOCMD) run main.go

format: 
	$(GOCMD) fmt ./...
	$(TEMPL) fmt .

tailwind:
	$(NPM) install -D tailwindcss
	$(NPX) tailwindcss -o ./static/public/tailwind.css

deploy: tailwind
	$(GOCMD) install github.com/a-h/templ/cmd/templ@latest
	$(TEMPL) generate
	$(GOCMD) build -tags netgo -ldflags '-s -w' -o app
	
