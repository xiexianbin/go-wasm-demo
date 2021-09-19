all: build_wasm cp_wasm_exec_js
	python3 -m http.server 9000

cp_wasm_exec_js:
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" static

build_wasm : main.go
	GO111MODULE=auto GOOS=js GOARCH=wasm go build -o static/main.wasm .
