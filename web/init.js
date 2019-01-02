
// __________________________________________________________________
//      Initialize Canvas
// ==================================================================

// purposely expose these global variables.
// It's easier to use them from WebAssembly code.

var canvas = document.querySelector('#game');
var ctx = canvas.getContext('2d');


// __________________________________________________________________
//      Loading Screen
// ==================================================================

function hideLoadingScreen() {
    document.querySelector('.loadbox').classList.add('hidden');
}


// __________________________________________________________________
//      Detect and Display Errors
// ==================================================================

// Mysterious errors seem to be happening a lot more often when webassembly
// is involved.  If you don't have a console (like phones), then this
// javascript will hopefully still pick them up, and display them
// "loudly" on the screen.

// TODO: add a way to send errors somewhere useful. (do HTTP post or something)



let errbox = document.querySelector('.errorbox');
let errclose = document.querySelector('.errclose');
let errlog = document.querySelector('#errlog');

function displayErrorMessage(err) {
    let line = document.createElement('div');
    line.classList.add('errline');
    line.innerText = err;
    if (line.innerText.length === 0) {
        line.innerText = "\n";
    }
    errlog.appendChild(line);
    errbox.classList.remove('hidden');
}

errclose.addEventListener('click', function(){
    errbox.classList.add('hidden');
});

// Redirect console log output to the the error display.
// https://stackoverflow.com/questions/11403107/capturing-javascript-console-log
(function(){
    var oldLog = console.log;
        console.log = function (message) {
        displayErrorMessage(message);
        oldLog.apply(console, arguments);
    };
})();


// copy all errors and display them on the visible screen.
window.addEventListener('error', function(event){
    displayErrorMessage(event.message);
});
