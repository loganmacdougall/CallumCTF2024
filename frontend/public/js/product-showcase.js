const images = Array.from(document.getElementsByClassName("product-image"))
const mini_images = Array.from(document.getElementsByClassName("product-mini-image"))

let active_index = 0

mini_images.forEach((mi, i) => {
  mi.addEventListener("click", () => {
    if (images[active_index].classList.contains("active")) {
      images[active_index].classList.remove("active")
    }
    if (mini_images[active_index].classList.contains("active")) {
      mini_images[active_index].classList.remove("active")
    }

    images[i].classList.add("active")
    mini_images[i].classList.add("active")

    active_index = i
  })
});

const clamp = (num, min, max) => {
  return Math.min(Math.max(num, min), max);
};

images.forEach((image, i) => {
  let mouseOver = false

  let moveImageFunc = (cx, cy) => {
    if (!image.classList.contains("active")) {
      return
    }

    let boundBox = image.getBoundingClientRect()

    let x = clamp(120 * (cx - boundBox.x) / boundBox.width - 10, 0, 100)
    let y = clamp(120 * (cy - boundBox.y) / boundBox.height - 10, 0, 100)

    image.style.objectPosition = `${x}% ${y}%`
    image.style.objectFit = `cover`
  }

  let resetImageFunc = () => {
    image.style.objectPosition = `center`
    image.style.objectFit = `scale-down`
  }

  image.addEventListener("mousemove", (e) => {
    moveImageFunc(e.x, e.y)
  })
  image.addEventListener("touchmove", (e) => {
    moveImageFunc(e.touches[0].clientX, e.touches[0].clientY)
  })
  image.addEventListener("mouseout", resetImageFunc)
  image.addEventListener("touchend", resetImageFunc)
})