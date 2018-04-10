PROJECT=app

clean:
	rm -Rf ./static/node_modules;
	rm -Rf ./tmp;
	rm -Rf ./build;

serve:
	fresh;

watch:
	cd ./static && webpack --config webpack.config.dev.js;

webpack:
	cd ./static && webpack

lint:
	cd ./static && npm run lint

test:
	go test;
	cd ./static && npm test;

install:
	cd ./server && go get -v ./...;
	cd ./server && go get github.com/pilu/fresh;
	cd ./static && npm install;

compile:
	GOOS=linux GOARCH=amd64 go build -o ./build/${PROJECT}-linux-amd64 .
	GOOS=darwin GOARCH=amd64 go build -o ./build/${PROJECT}-darwin-amd64 .
	cp -Rf ./templates ./build;

build: clean install webpack compile
	mkdir ./build/static;
	cp -Rf ./static/build ./build/static/build

