setup:
	cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./static

build:
	GOOS=js GOARCH=wasm go build -o static/main.wasm cmd/wasm/main.go

run:
	go run ./cmd/webserver/main.go