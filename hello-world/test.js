const fs = require('fs');
const buf = fs.readFileSync('./main.wasm');
require('./wasm_exec');
const go = new global.Go()

go.env = Object.assign({ TMPDIR: require("os").tmpdir() }, process.env);
(async () => { 
  await WebAssembly.instantiate(buf, go.importObject).then(result => {
    return go.run(result.instance);
  });
})()