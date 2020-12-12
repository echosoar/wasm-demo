const wasmFile = require('path').resolve(__dirname, './main.wasm');
const assert = require('assert');
require('../run')(wasmFile).then(() => {
  const obj = {
    name: 'test'
  }
  const result = global.callByThis.call(obj, 'name');
  assert(result === obj.name);
});