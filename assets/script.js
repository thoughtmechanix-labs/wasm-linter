const goWasm = new Go()

WebAssembly.instantiateStreaming(fetch("main.wasm"), goWasm.importObject)
    .then((result) => {
        goWasm.run(result.instance)

        document.getElementById("get-html").addEventListener("click" ,()=> {
            console.log("I am in the getMarkdownText")
            document.body.innerHTML += getMarkdownText(document.getElementById("markdown").value)
        })

    })