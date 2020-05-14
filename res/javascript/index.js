function scrollByID(id) {
    document.getElementById(id).scrollIntoView();
}

function scrollToBodyChild(i) {
    //document.body.children[i].scrollIntoView();
    window.scroll(0, window.innerHeight-70);
}

isFixed = false;

document.onscroll = function() {
    tabs = document.getElementById("navigationbar")
    if (!isFixed && document.body.scrollTop/window.innerHeight >= 0.14) {
        tabs.classList.add("fixedBar");

        isFixed = true;
    } else if (isFixed && document.body.scrollTop/window.innerHeight < 0.08) {
        tabs.classList.remove("fixedBar");

        isFixed = false;
    }
};