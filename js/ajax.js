const promise = new Promise((resolve, reject) => {
    //  variabel ajax bentuk http
    const xhr = new XMLHttpRequest()

    // untuk mengambil data dari server
    xhr.open("GET", "https://api.npoint.io/fe54a1a9dbc896357673", true)

    // untuk cek status nya
    xhr.onload = function () {
        if (xhr.status === 200) {
            resolve(JSON.parse(xhr.responseText))
            // JSON.parse = untuk nge parsing json kita
        } else if (xhr.status >= 400) {
            reject("error loading data")
        }
    }

    // untuk cek error nya
    xhr.onerror = function () {
        reject("network error")
    }


    //mengirim data
    xhr.send()
})


// promise.then((value) => {
//     console.log(value)
// }) .catch((reason) => {
//     console.log(reason)
// })


// ASYINC AWAIT
// untuk bekerja dengan promise dengan cara yang lebih nyaman, dipanggil “async/await”. Ini sangat mudah dipahami.jadi gunakan asyinc await biar promise d panggil dalam bentuk function

// variabel penampung for each
let testimoniaData = ""


async function getData() {
    try {
        const response = await promise
        console.log(response)
        // ini dari for each
        testimoniaData = response
        // ini func data html nya
        allTestimonial()

    } catch (err) {
        console.log(err)
    }
}
getData()

// func untuk mendaptkan data dari html nya
function allTestimonial() {
    let allTestimonialHTML = ""


    testimoniaData.forEach((card) => {
        allTestimonialHTML += `
    <div class="testimonial">
    <img class="profile-testimonial" src="${card.image}"/>
    <p class="quote">${card.quote}</p>
    <p class="author">-${card.author}</p>
    <p class="author">${card.rating}<i class="fa-regular fa-star"></i></p>
    </div>`
    })

    document.getElementById("testimonials").innerHTML = allTestimonialHTML
}


// func filter untuk mendapatkan rat data nya
function filterTestimonial(rating) {
    let filteredTestimonialHTML = ""

    const filteredData = testimoniaData.filter((card) => {
        return card.rating === rating
    })
    filteredData.forEach((card) => {
        filteredTestimonialHTML += `
        <div class="testimonial">
    <img class="profile-testimonial" src="${card.image}"/>
    <p class="quote">${card.quote}</p>
    <p class="author">-${card.author}</p>
    <p class="author">${card.rating}<i class="fa-regular fa-star"></i></p>
    </div>`
    })

    document.getElementById("testimonials").innerHTML = filteredTestimonialHTML
}

