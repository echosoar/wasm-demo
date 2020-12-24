const wasmFile = require('path').resolve(__dirname, './main.wasm');
const assert = require('assert');
require('../run')(wasmFile).then(() => {
  const result = global.typeTransform(
    123,
    'str',
    false,
    true,
    null,
    undefined,
    () => {
      return 'func'
    },
    NaN,
    [123, "name"],
    { num: 123, bool: false }
  );
  assert(result[0] === 123);
  assert(result[1] === 'str');
  assert(result[2] === false);
  assert(result[3] === true);
  assert(result[4] === null);
  assert(result[5] === undefined);
  assert(result[6] === 'func');
  assert(isNaN(result[7]));
  assert(Array.isArray(result[8]));
  assert(result[8][0] === 123);
  assert(result[8][1] === 'name');
  assert(result[9].num === 123);
  assert(result[9].bool === false);
});