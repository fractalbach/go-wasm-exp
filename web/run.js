
/* Define variables. These will get called by code in webassembly. */
var canvas = document.querySelector('#game');
var ctx = canvas.getContext('2d');

/* Initialize some canvas settings. (Totally optional) */
ctx.lineWidth = 4;
ctx.lineCap = 'round';


/** called when webassembly program finishes loading it's first canvas frame. */
function doneLoading() {
    document.querySelector('.loadbox').classList.add('hidden');
}

/** creates and shows a error overlay screen. */
function displayError(err) {
    let box = document.createElement('div');
    box.innerText = err;
    box.classList.add('errorbox');
    document.body.appendChild(box);
}

// window.addEventListener('load', ()=>{
    try {
        /* BEGIN: 3rd party code from golang repository: Jan 2019. */
        const go = new Go();
        WebAssembly.instantiateStreaming(fetch("web/main.wasm"), go.importObject).then((result) => {
            go.run(result.instance);
        });
        /* END: 3rd party code */
    } catch(err) {
        console.error(err);
        doneLoading();
    }
// });
