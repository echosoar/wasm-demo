const wasmFile = require('path').resolve(__dirname, './main.wasm');
const assert = require('assert');
require('../run')(wasmFile).then(() => {
  const obj = {
    name: 'test',
    age: 123
  }
  // changeObj(obj, doingType, key, targetValue)
  global.changeObj.call(obj, 'name', 'set', 123);
  global.changeObj.call(obj, 'age', 'delete');
  assert(obj.name === 123);
  assert(obj.age === undefined);
});