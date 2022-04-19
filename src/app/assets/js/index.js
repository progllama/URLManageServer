function handleDeleteClicked(userId, urlId) {
    sendRequest("/users/"+userId+"/urls/"+urlId, {"ID": parseInt(urlId)}, "DELETE")
}

function handleLogoutClicked(userId) {
  fetch("/logout", {
    method: 'DELETE',
    mode: 'cors',
    cache: 'no-cache', 
    credentials: 'same-origin',
    headers: {
      'Content-Type': 'application/json'
    },
    redirect: 'follow',
    referrerPolicy: 'no-referrer',
    body: {}
  })
}

function sendUrlUpdateForm() {
  sendUpdateForm("", "")
}

function sendUrlDeleteForm() {
  sendDeleteForm("", "")
}

function sendUpdateForm(url, body) {
  sendRequest(url, "DELETE", body)
}

function sendDeleteForm(url, body) {
  sendRequest(url, "PUT", body)
}

function sendForm(path, method, body) {
  
}

function sendRequest(url = '', data = {}, method='GET') {
  fetch(url, {
    method: method,
    mode: 'cors',
    cache: 'no-cache', 
    credentials: 'same-origin',
    headers: {
      'Content-Type': 'application/json'
    },
    redirect: 'follow',
    referrerPolicy: 'no-referrer',
    body: JSON.stringify(data)
  })
}

