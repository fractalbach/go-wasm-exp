/* ===================== Initialize Canvas ================= */

/** called when webassembly program finishes loading it's first canvas frame. */
function doneLoading() {
    document.querySelector('.loadbox').classList.add('hidden');
}

/* Define variables. These will get called by code in webassembly. */
var canvas = document.querySelector('#game');
var ctx = canvas.getContext('2d');

/* Initialize some canvas settings. (Totally optional) */
ctx.lineWidth = 4;
ctx.lineCap = 'round';



/* ===================== Run WebAssembly ================= */

window.addEventListener('load', function(){
    try {
        /* BEGIN: 3rd party code: golang repo: Jan 2019. */
        const go = new Go();
        WebAssembly.instantiateStreaming(fetch("web/main.wasm"), go.importObject).then((result) => {
            go.run(result.instance);
        });
        /* END: 3rd party code */
    } catch(err) {
        console.error(err);
        doneLoading();
    }
});
