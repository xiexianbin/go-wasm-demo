all: build_wasm
	npm run serve

build_wasm :
	GO111MODULE=auto GOOS=js GOARCH=wasm go build -o ./public/main.wasm ./wasm/main.go
