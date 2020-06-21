PROJECT=app

.PHONY: clean
clean:
	rm -Rf ./static/node_modules;
	rm -Rf ./tmp;
	rm -Rf ./build;

.PHONY: serve
serve:
	fresh;

.PHONY: watch-js
watch-js:
	cd ./static && webpack --config webpack.config.dev.js;

.PHONY: watch-go
watch-go:
	air -c air.conf

.PHONY: fmt
fmt:
	goimports -w pkg cmd internal
	gofmt -s -w pkg cmd internal

.PHONY: webpack
webpack:
	cd ./static && webpack

.PHONY: lint-go
lint-go:
	golangci-lint run

.PHONY: lint-js
lint-js:
	cd ./static && npm run lint

.PHONY: lint
lint: lint-go lint-js

.PHONY: test-js
test-js:
	cd ./static && npm test;

.PHONY: test-go
test-go:
	richgo test ./... -mod=readonly -v
	richgo test -v -race -coverpkg=./... -coverprofile=coverage.txt ./... -mod=readonly

.PHONY: test
test: test-js test-go

.PHONY: install-js
install-js:
	cd ./static && npm install;

.PHONY: install-go
install-go:
	go get -u -v ./...
	go mod tidy

.PHONY: install
install: install-js install-go

.PHONY: compile-
compile:
	GOOS=linux GOARCH=amd64 go build -a -o  ./build/${PROJECT}-linux-amd64 ./cmd/app/main.go
	# GOOS=darwin GOARCH=amd64 go build -a -o ./build/${PROJECT}-darwin-amd64 ./cmd/app/main.go
	cp -Rf ./templates ./build;

.PHONY: build
build: clean install webpack compile
	mkdir ./build/static;
	cp -Rf ./static/build ./build/static/build

.PHONY: local-setup
local-setup:
	/usr/bin/env ./.setup.sh
	go get -t github.com/kyoh86/richgo 
	go get -t golang.org/x/tools/cmd/goimports
	go get -t github.com/golangci/golangci-lint/cmd/golangci-lint
