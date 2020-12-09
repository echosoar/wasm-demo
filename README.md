# wasm-demo
Node.js 与 Golang 的各种 WebAssembly 示例

## Usage

进入到每一个示例的目录:

```bash
$ # 编译 golang 代码生成 wasm
$ GOOS=js GOARCH=wasm go build -o main.wasm
$
$ # 在 Node.js 中进行调用
$ node test.js
```

## Demo List
| 示例 | 效果 |
|---|---|
| [hello-world](./hello-world) | 控制台打印 `hello world` |
| [global-variable](./global-variable) | 往 js 的 `global` 上面挂载一个 **变量** |
| [global-function](./global-function) | 往 js 的 `global` 上面挂载一个 **方法** |

---

@echosoar License MIT