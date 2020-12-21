const wasmFile = require('path').resolve(__dirname, './main.wasm');
const assert = require('assert');
require('../run')(wasmFile).then(() => {
  class A {
    constructor(age) {
      this.age = age;
    }
    run() {
      return this.age;
    }
  }
  global.classA = A;
  const result = global.invokeClassRun();
  assert(result === 123);
});