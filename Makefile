setup:
	mv static/wasm_exec.js.disabled static/wasm_exec.js

re-setup:
	mv static/wasm_exec.js static/wasm_exec.js.disabled

build:
	GOOS=js GOARCH=wasm go build -o static/main.wasm cmd/wasm/main.go

run:
	GOOS=js GOARCH=wasm go build -o static/main.wasm cmd/wasm/main.go
	go run ./cmd/webserver/main.go

br:
	make build
	make run

git:
	mv static/wasm_exec.js static/wasm_exec.js.disabled
	git add .
	git commit -m "$m"
	git push