window.onload = () => {
    const listItems = document.getElementsByClassName("toggle-icon");
    for (let v of listItems) {
        v.addEventListener('click', function () {
            v.classList.toggle('fa-angle-down');
            v.classList.toggle('fa-angle-up');
        }, false);
    }

    const modalOpen = document.getElementsByClassName("modal-open")[0];
    const modalClose = document.getElementsByClassName("modal-close")[0];
    const modalForm = document.getElementsByClassName("modal-form")[0];

    let modalState = "CLOSE";
    modalOpen.addEventListener('click', () => {
        if (modalState == "OPEN") return;
        modalForm.classList.toggle("show");
        modalForm.classList.toggle("hide");
        modalState = "OPEN"
    })

    modalClose.addEventListener("click", () => {
        if (modalState == "CLOSE") return;
        modalForm.classList.toggle("show");
        modalForm.classList.toggle("hide");
        modalState = "CLOSE";
    });

    const submitButton = document.getElementsByClassName("submit")[0];
    submitButton.addEventListener("click", () => {
        const form = document.forms[0];

        const title = form.title.value;
        const url = form.url.value;
        const link = {
            title: title,
            url: url,
        }
        fetch(form.action, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(link)  // リクエスト本文に文字列化したJSON形式のデータを設定
        }).then((response) => {
            console.log(response)
        });
    })

    const linkList = document.getElementsByClassName("links")[0];
    fetch("/users/" + linkList.id + "/links", {
        method: "GET",
        headers: {
            'Content-Type': 'application/json'
        }
    }).then((response) => {
        const item = document.createElement("li");
        item.innerHTML = response
        linkList.appendChild(item);
    });
}