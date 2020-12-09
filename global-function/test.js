const wasmFile = require('path').resolve(__dirname, './main.wasm');
require('../run')(wasmFile).then(() => {
  console.log(global.addFun(1,2));
});