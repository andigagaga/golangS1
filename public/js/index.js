// console.log("andi ganteng")

let hamburgerIsOpen = false;
// hamburgerIsOpen == false
// !hamburgerIsOpen
function openHamburger() {
    let hamburgerNavContainer = document.getElementById("hamburger-nav-container");
    if (!hamburgerIsOpen) {
        hamburgerNavContainer.style.display = "block";
        hamburgerIsOpen = true
        console.log(hamburgerIsOpen)
    } else {
        hamburgerNavContainer.style.display = "none";
        hamburgerIsOpen = false
        console.log(hamburgerIsOpen)

    }
}

