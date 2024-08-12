const form = document.getElementById("login-form")
const usernameInput = document.getElementById("username-input")
const passwordInput = document.getElementById("password-input")
const errorMessageNode =  document.getElementById("error-message")

form.addEventListener('submit', async (e) => {
  e.preventDefault();

  const data = {
    "username": usernameInput.value,
    "password": passwordInput.value 
  }
  const urlData = new URLSearchParams(data).toString()

  let res
  try {
    res = await fetch("/login", {
      method: "POST",
      body: urlData,
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      }
    })
  } catch (err) {}

  if (!res) return
  if (res.status == 200) {
    return window.location.replace("/shop")
  }

  let errorMessage = await res.text()
  errorMessageNode.innerHTML = `Error: ${errorMessage}`
})