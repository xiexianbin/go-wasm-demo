all: build_wasm
	npm run serve

build_wasm :
	tinygo build -o ./public/main.wasm -target wasm ./wasm
