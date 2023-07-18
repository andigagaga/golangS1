// console.log("andi ganteng")
// enchapsulation pkai #

// class js
// class testimonial {
//     #quote = ""
//     #image = ""

//     constructor(quote, image) {
//         this.#quote = quote
//         this.#image = image
//     }
//      // pholimerphism
//      get quote() {
//         return this.#quote
//     }
//     get image() {
//         return this.#image
//     }
//     // ini dalam bentuk required harus ada authornya
//     get author() {
//         throw new Error("getAuthor() method must be implemented");
//       }


// }
// // inharitence = ntuk buat class di dalam class js  = dan disini kita jadikan parameter nya lebih fleksibel
// class authorTestimonials extends testimonial {
//     #author = ""


//     constructor(author, quote, image) {
//         super(quote, image)
//         this.#author = author
//     }

//     get author() {
//         return this.#author
//     }


// }
// class companyTestimonials extends testimonial {
//     #company = ""


//     constructor(company, quote, image) {
//         super(quote, image)
//         this.#company = company
//     }

//     get author() {
//         return this.#company + " Company";
//       }

// }

// const testimonial1 = new authorTestimonials(
// "Guswandi", 
// "Semangat Juang😘", 
// "https://images.unsplash.com/photo-1607604276583-eef5d076aa5f?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=774&q=80")

// const testimonial2 = new authorTestimonials (

// "Ituuu", 
// "Proggammer?🥲",
//  "https://images.unsplash.com/photo-1597851065532-055f97d12e47?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=870&q=80")

// const testimonial3 = new companyTestimonials (
// "Manusiaa", 
// "Gak Ngapa Ngapain😁", 
// "https://images.unsplash.com/photo-1615592389070-bcc97e05ad01?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=870&q=80")

// 
// // variabel off object untuk menampung semua data
// let testimoniaData = [testimonial1, testimonial2, testimonial3]

// // untuk menampung object d html nya
// let testimoniaHTML = ""

// // looping untuk store setiap data html nya...
// for(let i = 0; i < testimoniaData.length;i++) {
//     testimoniaHTML += `
//     <div class="testimonial">
//             <img class="profile-testimonial" src="${testimoniaData[i].image}"/>
//             <p class="quote">${testimoniaData[i].quote}</p>
//             <p class="author">-${testimoniaData[i].author}</p>
//         </div>`
// }



// // html dom nya agar data tampil d html nya yg sudah berdasarkan id d html
// document.getElementById("testimonials").innerHTML = testimoniaHTML



// dengan HOF DAN CALLBACK / ARAY OF OBJECT................................!!!!!!!!!!!!!!!!!!

const testimoniaData = [
    {
        author: "Guswandi",
        quote: "Semangat Juang😘",
        image: "https://media.istockphoto.com/id/1278947367/id/foto/keamanan-tato-selama-krisis-virus-corona.jpg?s=1024x1024&w=is&k=20&c=jsVQdi5wJMiNi73yZWJl5g2MvjfiEc4RjICPbLSMbVQ=",
        rating: 5
    },
    {
        author: "Genjer Sunowo",
        quote: "Proggammer?🥲",
        image: "https://images.unsplash.com/photo-1597851065532-055f97d12e47?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=870&q=80",
        rating: 4
    },
    {
        author: "Sratowo",
        quote: "Gak Ngapa Ngapain😁",
        image: "https://images.unsplash.com/photo-1615592389070-bcc97e05ad01?ixlib=rb-4.0.3ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=870q=80",
        rating: 3
    },
    {
        author: "MegaChan",
        quote: "Pemersatu Bangsa🥲",
        image: "https://images.unsplash.com/photo-1687732589937-4f29aa3c82dc?ixlib=rb-4.0.3ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=870q=80",
        rating: 3
    },
    {
        author: "Bang Joko",
        quote: "Jangan Pergi🤣",
        image: "https://images.unsplash.com/photo-1688240371753-24368f363a47?ixlib=rb-4.0.3ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=871q=80",
        rating: 1
    },

];

// let testimonialHTML = ""

// // pake FOREACH / KARENA TIDAK PERLU MENGEMBALIKAN NILAINYA
testimoniaData.forEach((card) => {
    testimonialHTML += `
    <div class="testimonial">
    <img class="profile-testimonial" src="${card.image}"/>
    <p class="quote">${card.quote}</p>
    <p class="author">-${card.author}</p>
    <p class="author">${card.rating}<i class="fa-regular fa-star"></i></p>
    </div>`
})

document.getElementById("testimonials").innerHTML = testimonialHTML



// FUNC UNTUK RAT TESTIMONIAL DALAM BENTUK FILTER AGAR BISA MENYARING DATA

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


// func untuk rat testimonial secara defaultnya
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

// pemanggilan untuk secara default/eksekusi awalnya
allTestimonial()
