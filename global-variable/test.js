const wasmFile = require('path').resolve(__dirname, './main.wasm');
require('../run')(wasmFile, true).then(() => {
  console.log(global.thisIsOneTwoThree, global.thisIsOneTwoThree === 123);
});