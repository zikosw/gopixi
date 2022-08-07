GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o ../public/main.wasm ./cmd
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ../src/