#!/bin/bash

pwd
rm -rf dist

# copy repo to dist folder
# cp -r . ./dist/ --exclude ./dist
rsync -av --progress . ./dist --exclude ./dist

# compile wasm file
GOOS=js GOARCH=wasm go build -o ./dist/testgame.wasm testgame

# install wasmserve
go get github.com/hajimehoshi/wasmserve
# copy js helper from golang
cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./dist/

# make an index.html
echo "<!DOCTYPE html>
<script src="wasm_exec.js"></script>
<script>
// Polyfill
if (!WebAssembly.instantiateStreaming) {
  WebAssembly.instantiateStreaming = async (resp, importObject) => {
    const source = await (await resp).arrayBuffer();
    return await WebAssembly.instantiate(source, importObject);
  };
}

const go = new Go();
WebAssembly.instantiateStreaming(fetch(\"testgame.wasm\"), go.importObject).then(result => {
  go.run(result.instance);
});
</script>" >> ./dist/index.html
cd dist

clear
echo "Serving up wasm on http://localhost:8080/"

# serve it up
wasmserve .
