all: build_wasm
	quasar dev

build_wasm :
	GO111MODULE=auto GOOS=js GOARCH=wasm go build -o ./public/main.wasm ./wasm/main.go
