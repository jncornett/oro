hello.wasm: cmd/hello/*.go
	env GOOS=js GOARCH=wasm go build -o hello.wasm ./cmd/hello

default: hello.wasm
