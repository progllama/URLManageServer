
window.onload = () => {
  console.log("Loading javascript.")

  const logout = document.getElementById("logout");
  if (logout != null){
    logout.addEventListener("click", sendLogout);
  }

  const pathName = location.pathname.split("/")
  const userIdPos = pathName.indexOf("users")+1
  const userId = pathName[userIdPos]
  const deletionBtns = document.querySelectorAll(".delete_url")
  let counter = 0;
  for (let i of deletionBtns){
    console.log("success to get delete.")
    const id = counter
    i.addEventListener("click", () => { sendUrlDeleteRequest(userId,i.id.split("_")[1], "url_" + id) })
    counter += 1;
  }
};


function handleDeleteClicked(userId, urlId) {
    sendRequest("/users/"+userId+"/urls/"+urlId, {"ID": parseInt(urlId)}, "DELETE")
}

function sendLogout(event) {
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
  .then(() => alert("Success."))
  .catch(() => alert("Failed."))
}

function sendUrlUpdateForm(userId, urlId) {
  return (event) => {
    const form = new FormData()
    const inputNames = [
      'title',
      'url',
      'description',
      'note'
    ]
    for (let name of inputNames)
    {
      const input = document.querySelector('input[name=' + name + ']');
      form.append(name, input.value)
    }
  
    fetch( '/users/'+userId+'/urls'+urlId, {
      method: 'PUT',
      body: form
    })
    .then(res => res.json())
    .catch(err => console.log(err))
  } 
}

function sendUrlDeleteRequest(userId, urlId, itemId) {
  const url = '/users/'+userId.toString()+'/urls/' + urlId.toString()
  fetch(
    url,
    {
      method: 'DELETE',
    }
  )
  .then(res => {
    if (!res.ok) {
      alert("削除に失敗しました。");
    };
    alert("削除に成功しました。")
    console.log(itemId)
    element = document.getElementById(itemId)
    element.remove()
  })
  .catch(err => console.log(err))
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

