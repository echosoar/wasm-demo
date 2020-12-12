const wasmFile = require('path').resolve(__dirname, './main.wasm');
const assert = require('assert');
require('../run')(wasmFile).then(() => {
  global.arrayData = [];
  global.changeArray('push', 123);
  assert(global.arrayData.length === 1);
  assert(global.arrayData[0] === 123);
  const pooValue = global.changeArray('pop');
  assert(global.arrayData.length === 0);
  assert(pooValue === 123);
});