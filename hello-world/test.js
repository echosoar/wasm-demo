const wasmFile = require('path').resolve(__dirname, './main.wasm');
require('../run')(wasmFile, true);