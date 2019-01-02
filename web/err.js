
/* ===================== Display Errors on Screen ================= */
/*
Mysterious errors seem to be happening a lot more often when webassembly
is involved.  If you don't have a console (like phones), then this javascript
will hopefully still pick them up, and display them "loudly" on the screen.
*/

let errbox = document.querySelector('.errorbox');
let errclose = document.querySelector('.errclose')

function displayError(err, info) {
    let line = document.createElement('div');
    if (info) {
        line.innerText = "[Info]: " + err;
        line.classList.add('infoline');
    } else {
        line.innerText = "[Error]: " + err;
    }
    line.classList.add('errline');
    errbox.appendChild(line);
    errbox.classList.remove('hidden');
}

errclose.addEventListener('click', function(){
    errbox.classList.add('hidden');
});

(function(){
    var oldLog = console.log;
    console.log = function (message) {
        displayError(message, 1);
        oldLog.apply(console, arguments);
    };
})();


/* copy all errors and display them on the visible screen. */
window.addEventListener('error', function(event){
    displayError(event.message);
});
