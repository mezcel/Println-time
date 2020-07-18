// Button Key bindings
document.onkeyup = function(e){
    switch(e.which) {
        case 39:        // rt arrow
        case 76:        // l
            document.getElementById( "next" ).click()
            break;
        case 72:        // h
        case 37:        // lt arrow
            document.getElementById( "back" ).click()
            break;
    }
}
