window.onload = () => {
    const listItems = document.getElementsByClassName("toggle-icon");
    for (let v of listItems) {
        v.addEventListener('click', function () {
            v.classList.toggle('fa-angle-down');
            v.classList.toggle('fa-angle-up');
        }, false);
    }
}