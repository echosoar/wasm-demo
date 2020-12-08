# wasm-demo
Golang webassembly demo

```bash
$ GOOS=js GOARCH=wasm go build -o main.wasm
```
```bash
$ cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./
```

