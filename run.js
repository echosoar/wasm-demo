const fs = require('fs');
const { execSync } = require('child_process');
const wasmExec = execSync('go env GOROOT').toString().replace(/\n/g, '') + '/misc/wasm/wasm_exec.js';
require(wasmExec);

// isWait 代表是否要等到go代码执行结束之后再执行js代码
// 如果存在js调用go方法，那就不能等
module.exports = async (file, isWait) => {
  const buf = fs.readFileSync(file);
  const go = new global.Go();
  return WebAssembly.instantiate(buf, go.importObject).then(result => {
    if (isWait) {
      // 这个await执行完成代码go代码已经执行结束退出了
      return go.run(result.instance);
    } else {
      go.run(result.instance);
    }
    return;
  });
}
