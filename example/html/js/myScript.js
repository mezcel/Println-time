/* myScript.js */

// Toggle w3css dropdown menu display
function toggleDropdown() {
    var x = document.getElementById("Menu");

    if (x.className.indexOf("w3-show") == -1) { 
        x.className += " w3-show";
    } else {
        x.className = x.className.replace(" w3-show", "");
    }
}

// Button Key bindings
document.onkeyup = function(e){
    var keyPress = e.which;
    switch(keyPress) {
        case 39:        // rt arrow
        case 76:        // l
            // navigate to next position
            var btnNext = document.getElementById( 'btnNext' );

            if ( btnNext ) {
                btnNext.click();
            }
            break;
        case 72:        // h
        case 37:        // lt arrow
            // navigate to previous position
            var btnBack = document.getElementById( 'btnBack' );
            var btnIndex = document.getElementById( 'btnIndex' );

            if ( btnBack ) {
                btnBack.click();
            }
            if ( btnIndex ) {
                btnIndex.click();
            }
            break;
        case 32:        // spacebar
            // toggle dropdown mwnu
            toggleDropdown();
            break;
        default:
            console.log("keyPress:",keyPress);
    }
}

document.addEventListener("DOMContentLoaded", function() {
    console.log("Page loaded ...");
    /*var btnMenu = document.getElementById( 'btnMenu' );
    if (btnMenu) {
        btnMenu.focus();
    }*/
});
