function handleDeleteClicked(userId, urlId) {
    sendRequest("/users/"+userId+"/urls/"+urlId, {}, "DELETE")
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