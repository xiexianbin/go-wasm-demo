all: build_wasm
	quasar dev

build_wasm :
	tinygo build -o ./public/main.wasm -target wasm ./wasm
