
import './wasm_exec.js'


export default function(job) {
    console.log("init gowasm")
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
        go.run(result.instance);
        console.log("gowasm has run")
        job()
    });
}