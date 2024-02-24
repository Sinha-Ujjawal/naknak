// -----------------------Loading WASM---------------------------
const go = new Go(); // Defined in wasm_exec.js
const WASM_URL = 'main.wasm';

let wasm_promise;

if ('instantiateStreaming' in WebAssembly) {
    wasm_promise = WebAssembly
        .instantiateStreaming(fetch(WASM_URL), go.importObject)
        .then(obj => {
            let wasm = obj.instance;
            go.run(wasm);
            return wasm;
        });
} else {
    wasm_promise = fetch(WASM_URL)
        .then(resp => resp.arrayBuffer())
        .then(bytes => WebAssembly
            .instantiate(bytes, go.importObject)
            .then(obj => {
                let wasm = obj.instance;
                go.run(wasm);
                return wasm;
            })
        );
}

// -----------------------UI Setup-----------------------
const textToEncode = document.getElementById("textToEncode");
const textToDecode = document.getElementById("textToDecode");

if (textToEncode !== undefined) {
    let timeout;
    textToEncode.addEventListener("keydown", (_event) => {
        clearTimeout(timeout);
        textToDecode.disabled = true;
        timeout = setTimeout(
            () => wasm_promise.then(_ => {
                textToDecode.value = encode(textToEncode.value);
                textToDecode.disabled = false;
            }),
            500
        ); 
    });
} else {
    console.error("id `textToEncode` not found!");
}

if (textToDecode !== undefined) {
    let timeout;
    textToDecode.addEventListener("keydown", (_event) => {
        clearTimeout(timeout);
        textToEncode.disabled = true;
        timeout = setTimeout(
            () => wasm_promise.then(_ => {
                textToEncode.value = decode(textToDecode.value);
                textToEncode.disabled = false;
            }),
            500
        ); 
    });
} else {
    console.error("id `textToDecode` not found!");
}
