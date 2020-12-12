const wasmFile = require('path').resolve(__dirname, './main.wasm');
const assert = require('assert');
require('../run')(wasmFile).then(() => {
  const date = new Date();
  global.testFun = () => {
    return date;
  }
  global.testFun2 = (arg) => {
    return arg;
  }
  const result = global.callFun('testFun', undefined);
  const result2 = global.callFun2('testFun2', 123);
  assert(typeof result === 'object');
  assert(typeof result2 === 'number');
  assert(+result === +date);
  assert(result2 === 123);
});