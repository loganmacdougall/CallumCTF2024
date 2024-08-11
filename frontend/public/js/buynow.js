async function buyNow(id) {
  const data = {"id": id}
  const urlData = new URLSearchParams(data).toString()

  let res
  try {
    res = await fetch(`/buynow`, {
      method: "POST",
      body: urlData,
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      }
    })
  } catch (err) {}

  if (!res) return
  if (res.status == 200) {
    window.location.replace(`/thankyou?id=${id}`)
  } else if (res.status == 302) {
    let link = await res.text()
    window.location.replace(link)
  } else if (res.status != 400) {
    window.location.replace(`/login`)
  }
}