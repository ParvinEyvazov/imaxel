mkfile_dir := $(dir $(mkfile_path))

setup:
	mv D:\Projects\imaxel\static\wasm_exec.js.disabled D:\Projects\imaxel\static\wasm_exec.js

re-setup:
	mv D:\Projects\imaxel\static\wasm_exec.js D:\Projects\imaxel\static\wasm_exec.js.disabled

build:
	GOOS=js GOARCH=wasm go build -o static/main.wasm cmd/wasm/main.go

run:
	GOOS=js GOARCH=wasm go build -o static/main.wasm cmd/wasm/main.go
	go run ./cmd/webserver/main.go

br:
	make build
	make run

git:
	git add .
	git commit -m "$m"
	git push