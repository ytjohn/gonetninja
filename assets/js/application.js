require("expose-loader?exposes=$,jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require("@fortawesome/fontawesome-free/js/all.js");
require("jquery-ujs/src/rails.js");
const {hide} = require("@popperjs/core");

$(() => {

});

let editments = document.getElementsByClassName("editelement")
let queryParams = new URLSearchParams(window.location.search);

function toggleeditmode() {
    for (let i = 0; i < editments.length; i++) {
        editments[i].classList.toggle("edithidden")
    }
}

function hideeditments() {
    console.log("hide edit mode")
    for (let i = 0; i < editments.length; i++) {
        editments[i].classList.add("edithidden")
    }
}

function showeditments() {
    console.log("edit mode enabled");
    for (let i = 0; i < editments.length; i++) {
        editments[i].classList.remove("edithidden")
    }
}

function initialeditmode() {
    editmode = queryParams.get("editmode",);
    if (editmode === null) { editmode = "false"; }
    if (JSON.parse(editmode.toLowerCase())) {
        showeditments();
    } else {
        hideeditments();
    }
}

document.addEventListener("load", initialeditmode());
document.getElementById('toggleeditor').addEventListener('click', toggleeditmode); // add an event listener to the button