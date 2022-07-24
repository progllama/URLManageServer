// const UIComponents = {
//     currentActiveSidebarItem: null,
// }

const registry = {
    DeleteConfirmModal: null,
    ErrMsgState: "HIDE"
}

const LinkDeleteURL = "/users"

const HIDE = "HIDE";
const SHOW = "SHOW";
const OPEN = "OPEN";
const CLOSE = "CLOSE";

let _sidebar;
function setSidebar(s) {
    _sidebar = _sidebar ?? s;
}

function getSidebar() {
    return _sidebar;
}

let _currentActiveSidebarMenu;
function setSidebarMenuActive(newItem) {
    setSidebarMenuInactive()
    if (_currentActiveSidebarMenu == newItem) { return };
    const title = document.getElementById("sidebar-header-text");
    _currentActiveSidebarMenu = newItem;
    _currentActiveSidebarMenu.classList.toggle("active");
    title.innerHTML = newItem.innerHTML;
}

function setSidebarMenuInactive() {
    if (_currentActiveSidebarMenu == null) { return };
    _currentActiveSidebarMenu.classList.toggle("active");
}

function getSidebarMenuActive() {
    return _currentActiveSidebarMenu;
}

function setDeleteConfirmModal(m) {
    registry.DeleteConfirmModal = m;
}

function getDeleteConfirmModal() {
    return registry.DeleteConfirmModal;
}

window.onload = () => {
    setSidebar(document.getElementsByClassName("sidebar-wrapper")[0]);

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
            renderPanels();
        }).catch((err) => {
            console.log(err);
            const msg = document.getElementsByClassName("fault-message")[0];
            msg.classList.toggle("hide")
            msg.classList.toggle("show")
            registry.ErrMsgState = "SHOW"
        });

        document.body.addEventListener("click", () => {
            if (registry.ErrMsgState == "HIDE") return;
            const msg = document.getElementsByClassName("fault-message")[0];
            msg.classList.toggle("hide")
            msg.classList.toggle("show")
            registry.ErrMsgState = "HIDE";
        });
    })

    renderPanels();

    const deleteConfirmModal = new DeleteConfirmModal(document.getElementsByClassName("delete-confirm-modal")[0]);
    deleteConfirmModal.init();
    setDeleteConfirmModal(deleteConfirmModal);

    const newListButton = document.getElementsByClassName("add-dir")[0];
    newListButton.addEventListener("click", () => {
        const sidebar = getSidebar();
        const input = document.createElement("input");
        sidebar.appendChild(input);
        input.focus();
        input.addEventListener("change", (event) => {
            sidebar.removeChild(input);
            const wrapper = document.createElement("div");
            const arrow = document.createElement("i");
            const text = document.createElement("span");
            text.innerHTML = event.target.value;
            arrow.classList.add("fa-solid");
            arrow.classList.add("toggle-icon")
            arrow.classList.add("fa-angle-down")
            fetch("/users/" + UserID + "/link_lists/", {
                method: "POST",
                body: JSON.stringify({
                    title: event.target.value
                })
            }).then(response => {
                return response.json();
            }).then(body => {
                console.log(body);
                wrapper.draggable = true;
                wrapper.classList.add("sidebar-item")
                wrapper.id = "side-" + body.ID;
                wrapper.appendChild(arrow)
                wrapper.appendChild(text);
                sidebar.appendChild(wrapper);
                arrow.addEventListener('click', function () {
                    arrow.classList.toggle('fa-angle-down');
                    arrow.classList.toggle('fa-angle-up');
                });
            })
        })
    });

    // サイドバーのタイトル部の初期化
    const sidebarTitles = document.getElementsByClassName("sidebar-item-title");
    for (let item of sidebarTitles) {
        item.addEventListener("click", () => {
            setSidebarMenuActive(item);
        })
    }

    // サイドバーのリストのDELETE処理
    // サイドメニューの要素を全て取得
    const sidebarItems = document.getElementsByClassName("sidebar-item");
    for (let item of sidebarItems) {
        item.addEventListener("dragstart", (event) => {
            event.dataTransfer.setData("text/plain", event.target.id)
        })
    }
    // ゴミ箱にドロップした時のイベントをハンドリング
    const sidebarTrash = document.getElementsByClassName("remove-dir")[0];
    sidebarTrash.addEventListener("dragover", (event) => {
        // dropを有効化
        event.preventDefault();
    })
    sidebarTrash.addEventListener("drop", (event) => {
        const target = document.getElementById(event.dataTransfer.getData("text/plain"));
        fetch("/users/" + UserID + "/link_lists/" + event.dataTransfer.getData("text/plain").slice(5), {
            method: "DELETE",
        }).then(response => {
            return response.json();
        }).then(status => {
            const sidebar = document.getElementsByClassName("sidebar-wrapper")[0];
            sidebar.removeChild(target)
        });
    });
}

function buildPanel(link) {
    const panel = document.createElement("div");
    panel.classList.add("link-panel");

    const summary = document.createElement("div");
    summary.classList.add("summary");
    summary.classList.add("panel-item");

    const title = document.createElement("div");
    title.innerHTML = "title : " + link.title;

    const url = document.createElement("div");
    const anchor = document.createElement("a");
    anchor.href = link.url;
    let suffix = "";
    if (link.url.length > 50) { suffix = "..." }
    anchor.innerHTML = "url : " + link.url.substr(0, 50) + suffix;
    anchor.target = "_blank";
    anchor.rel = "noopener noreferrer";
    url.appendChild(anchor);

    summary.appendChild(title);
    summary.appendChild(url);

    panel.appendChild(summary);

    const panelButtons = document.createElement("div");
    panelButtons.classList.add("panel-buttons");
    const detailButton = buildDetailButton();
    const editButton = buildEditButton();
    const deleteButton = buildDeleteButton(link.title, UserID, link.ID);
    panelButtons.appendChild(detailButton);
    panelButtons.appendChild(editButton);
    panelButtons.appendChild(deleteButton);
    panel.appendChild(panelButtons);

    return panel;
}

function buildDetailButton() {
    const detailButton = document.createElement("div");
    detailButton.classList.add("panel-item");
    const detailIcon = document.createElement("i");
    detailIcon.classList.add("fa");
    detailIcon.classList.add("fa-list-alt");
    detailButton.appendChild(detailIcon);

    detailButton.addEventListener("click", () => {
        console.log("detail");
    })

    return detailButton;
}

function buildEditButton() {
    const editButton = document.createElement("div");
    editButton.classList.add("panel-item");
    const editIcon = document.createElement("i");
    editIcon.classList.add("fa");
    editIcon.classList.add("fa-edit");
    editButton.appendChild(editIcon);

    editButton.addEventListener("click", () => {
        console.log("edit");
    })

    return editButton;
}

function buildDeleteButton(title, uid, id) {
    const deleteButton = document.createElement("div");
    const trashIcon = document.createElement("i")
    trashIcon.classList.add("fa");
    trashIcon.classList.add("fa-trash");
    trashIcon.ariaHidden = true;
    deleteButton.appendChild(trashIcon);
    deleteButton.classList.add("panel-item");

    deleteButton.addEventListener("click", () => {
        const modal = getDeleteConfirmModal();
        modal.show(title, uid, id);
    })

    return deleteButton;
}

class DeleteConfirmModal {
    constructor(elem) {
        this.elem = elem
        this.active = false;
    }
    init() {
        const cancelButton = document.getElementsByClassName("delete-cancel")[0];
        cancelButton.addEventListener("click", () => {
            this.cleanup();
            this.hide();
        })
    }
    show(title, uid, id) {
        const placeholder = document.getElementsByClassName("delete-target")[0];
        placeholder.innerHTML = title;
        const deleteButton = document.getElementsByClassName("delete-execute")[0];
        // FIX
        deleteButton.addEventListener("click", () => {
            fetch(LinkDeleteURL + "/" + uid + "/links/" + id, { method: "DELETE" })
                .then((response) => { renderPanels() });
            this.hide();
        })
        this.elem.style.display = "block"
        this.active = true;
    }

    hide() {
        this.elem.style.display = "none";
        this.active = false;
    }
    cleanup() {
        const placeholder = document.getElementsByClassName("delete-target")[0];
        placeholder.innerHTML = "";
    }
}

function renderPanels() {
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
            const panel = buildPanel(v);
            linkList.appendChild(panel);
        }
    })
}