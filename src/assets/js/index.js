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
        const form = document.forms[0];
        const title = form.title.value;
        const url = form.url.value;
        form.title.value = "";
        form.url.value = "";
        modalState = "CLOSE";
    });

    const submitButton = document.getElementsByClassName("submit")[0];
    submitButton.addEventListener("click", () => {
        const form = document.forms[0];
        if (modalState == "OPEN") {
            modalForm.classList.toggle("show");
            modalForm.classList.toggle("hide");
            modalState = "CLOSE";
        }

        const title = form.title.value;
        const url = form.url.value;
        form.title.value = "";
        form.url.value = "";
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
            const linkList = document.getElementsByClassName("links")[0];
            while (linkList.firstChild) {
                linkList.removeChild(linkList.firstChild);
            }

            fetch("/users/" + linkList.id + "/links", {
                method: "GET",
                headers: {
                    'Content-Type': 'application/json'
                }
            }).then((response) => {
                return response.json();
            }).then((data) => {
                for (let v of data.links) {
                    const item = document.createElement("li");
                    item.innerHTML = v.title + " " + v.url
                    linkList.appendChild(item)
                }
            })
        });
    })

    const linkList = document.getElementsByClassName("links")[0];
    fetch("/users/" + linkList.id + "/links", {
        method: "GET",
        headers: {
            'Content-Type': 'application/json'
        }
    }).then((response) => {
        return response.json();
    }).then((data) => {
        for (let v of data.links) {
            const item = document.createElement("li");
            item.innerHTML = v.title + " " + v.url
            linkList.appendChild(item)
        }
    })
}