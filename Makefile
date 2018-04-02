PROJECT=app

clean:
	rm -Rf ./static/node_modules;
	rm -Rf ./tmp;
	rm -Rf ./bin;

serve:
	fresh;

watch:
	cd ./static && webpack --config webpack.config.dev.js

install:
	go get -v ./...
	go get github.com/pilu/fresh
	cd static && npm install;

compile:
	GOOS=linux GOARCH=amd64 go build -o bin/${PROJECT}-linux-amd64;
	GOOS=darwin GOARCH=amd64 go build -o bin/${PROJECT}-darwin-amd64;

build:
	clean
	install
	compile

