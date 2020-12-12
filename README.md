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
| [1. hello-world](./hello-world) | 控制台打印 `hello world` |
| [2. global-variable](./global-variable) | 往 js 的 `global` 上面挂载一个 **变量** |
| [3. global-function](./global-function) | 往 js 的 `global` 上面挂载一个 **方法** |
| [4. call-by-this](./call-by-this) | go 中通过 `this` 获取 js 的当前上下文 |
| [5. call-function](./call-function) | go 中调用 js `方法` |
| [6. change-object](./change-object) | go 中修改 js `Object` |
| [7. change-array](./change-array) | go 中修改 js `Array` |

---

©  [LICENSE](./LICENSE) MIT by [echosoar](https://github.com/echosoar)