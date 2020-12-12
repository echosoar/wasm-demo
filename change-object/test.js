const wasmFile = require('path').resolve(__dirname, './main.wasm');
const assert = require('assert');
require('../run')(wasmFile).then(() => {
  const obj = {
    name: 'test'
  }
  // changeObj(obj, doingType, key, targetValue)
  global.changeObj.call(obj, 'name', 123);
  assert(obj.name === 123);
});