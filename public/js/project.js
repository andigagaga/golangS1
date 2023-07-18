// array
// let namaSiswa = ["guswandi", "wahyu", "ira"];
// console.log(namaSiswa);
// console.log(namaSiswa[0]);
// console.log(namaSiswa[1])

// let arrayName = ["situs", "web","wibu"]
// console.log(arrayName)
// console.log(arrayName[2])

// let calonSiswa = ["bayu", "agus", "asep"];
// console.log(calonSiswa);
// console.log(calonSiswa[2])

// object
// let personal1 = {
//     name: "Guswandi",
//     alamat: "padang",
//     umur: 19,
// };
// let personal2 = {
//     name: "candra",
//     alamat: "padang",
//     umur: 18,
// }
// console.log(personal1);
// console.log(personal2);
// console.log(personal1.alamat);
// console.log(personal2.umur);
// let dokter = {
//     name: "babi",
//     alamat: "joko tingir",
//     umur: 126,
// }
// console.log(dokter)
// console.log(dokter.alamat)
// console.log(dokter.umur)

// array of object
// let dataWibu = [
//     {name: "andi", alamat: "padang"},
//     {name: "bagong", alamat: "jambi"}
// ];
// console.log(dataWibu[0]);


// variabel dataProject nya array of object
let dataProject = []
// alert("gw ganteng tauu")

// fungsi dari add projectnya
function addProject(event) {
    event.preventDefault();

    // variabel dari data add projectnya
    let title = document.getElementById("input-project-title").value;
    let stardate = document.getElementById("input-startdate").value;
    let enddate = document.getElementById("input-enddate").value;
    let description = document.getElementById("input-description").value;
    let image = document.getElementById("input-project-image").files;


    // variabel untuk menghitung jarak dursi pda project
    let start = new Date(stardate)
    let end = new Date(enddate)

    let timeDistance = end - start
    console.log(timeDistance)

    let distanceSecond = Math.floor(timeDistance / 1000)
    let distanceMinutes = Math.floor(distanceSecond / 60)
    let distanceHours = Math.floor(distanceMinutes / 60)
    let distanceDays = Math.floor(distanceHours / 24)
    let distanceWeeks = Math.floor(distanceDays / 7)
    let distanceMonths = Math.floor(distanceWeeks / 4)
    let distanceYears = Math.floor(distanceMonths / 12)
    // floor -> 1.8 -> 1
    // ceil -> 1.4 -> 2
    // round -> 1.3 -> 1

    let distance = "";

    // kondision untuk jarak waktu project
    if (distanceDays < 7) {
        distance = distanceDays + "hari";
    } else if (distanceWeeks < 4) {
        distance = distanceWeeks + "minggu";
    } else if (distanceMonths < 12) {
        distance = distanceMonths + "bulan";
    } else {
        distance = distanceYears + "tahun";
    }

    // let start = new Date(stardate)
    // let end = new Date(enddate)
    // let selisih = start - end
    // let days = selisih / (1000 * 60 * 60 * 24)
    // let weeks = Math.floor(days/7)
    // let month = Math.floor(weeks/4)
    // let  year= Math.floor(month/12)
    // let durasi = ""

    // if(days > 0) {
    //     durasi = days + "hati"
    // }else if(weeks > 0) {
    //     durasi = weeks + "minggu"
    // }else if(month > 0) {
    //     durasi = month + "bulan"
    // }else if(year > 0) {
    //     durasi = year + "tahun"
    // }
    // else {
    //     distance= `${distanceMonth} + "bulan"`
    // };







    // variabel untuk menghitung selisih waktu post nya
    // let mulai = new Date(stardate);
    // let akhir = new Date(enddate);
    // let selisih = akhir.getTime() - mulai.getTime();
    // let durasi1 = Math.floor(selisih / (1000 * 60 * 60 * 24));
    // let durasi2 = Math.floor(selisih / (1000 * 60 * 60 * 24 *30));
    // console.log(selisih)
    // if (durasi1 >= 30) {
    //     return `${durasi2} Bulan`
    // };



    // variabel untuk data pada iconnya
    const nodejsIcon = '<i class="fa-brands fa-node-js"></i>';
    const reactjsIcon = '<i class="fa-brands fa-react"></i>';
    const nextjsIcon = '<i class="fa-brands fa-android"></i>';
    const typescriptIcon = '<i class="fa-brands fa-java"></i>';

    // variabel untuk dapatkan value dan data dari chekbox nya
    let cbnodejs = document.getElementById("nodejs").checked ? nodejsIcon : "";
    let cbReactjs = document.getElementById("reactjs").checked ? reactjsIcon : "";
    let cbnextjs = document.getElementById("nextjs").checked ? nextjsIcon : "";
    let cbtypescript = document.getElementById("typescript").checked ? typescriptIcon : "";

    // untuk membuat url dari image nya
    image = URL.createObjectURL(image[0]);
    console.log(image);


    // untuk membuat object dari data projectnya
    let project = {
        title,
        distance,
        cbnodejs,
        cbReactjs,
        cbnextjs,
        cbtypescript,
        description,
        image,
        postAt: new Date(),
    }
    // 

    dataProject.push(project);
    console.log(project)
    console.log(dataProject)

    renderProject();


    // ini get untuk mengosongkan data setelah kita input/submit data
    document.getElementById("input-project-title").value = "";
    document.getElementById("input-startdate").value = "";
    document.getElementById("input-enddate").value = "";
    document.getElementById("input-description").value = "";
    document.getElementById("nodejs").checked = false;
    document.getElementById("nextjs").checked = false;
    document.getElementById("reactjs").checked = false;
    document.getElementById("typescript").checked = false;
    document.getElementById("input-project-image").value = "";

}


// fungsi untuk menampilkan data add project kita ke halaman html/html
function renderProject() {
    document.getElementById("mockup").innerHTML = "";
    // fungsion render itu biar kita mnggil ulang data dari form project .
    // looping
    for (let index = 0; index < dataProject.length; index++) {
        document.getElementById("mockup").innerHTML += `
            
            <div class="project-items1" id="project-items1">
                <div class="project-items-container">
                    <div class="project-items-image">
                        <img src="${dataProject[index].image}" alt="project-list1"  />
                        <a href="project-detail.html" target="_blank">
                            <button>Detail Project</button>
                        </a>
                        
                        
                    </div>
                    <div class="project-items-judul">
                        <h2>${dataProject[index].title}</h2>
                        <h6>${getFullTime(dataProject[index].postAt)}</h6>
                        <h5 class="project-items-duration">durasi : ${dataProject[index].distance}</h5>
                        
                    </div>
                    <div class="project-list-paraf">
                        <p>${dataProject[index].description}</p>
                    </div>
                    <div class="project-list-icon" id="project-list-icon">
                        ${dataProject[index].cbnodejs}
                        ${dataProject[index].cbReactjs}
                        ${dataProject[index].cbnextjs}
                        ${dataProject[index].cbtypescript}
                    </div>
                    <div class="button-project" id="button-project">
                        <button class="button-project-edit">
                            <a href="#" >edit</a>
                        </button>
                        <button class="button-project-delete">
                            <a href="#">delete</a>
                        </button>
                    </div>
                </div>
            </div>
       
    `;
    }

}

// untuk menghitung waktu  saat post project

function getFullTime(time) {
    let bulan = ["Jan", "Feb", "March", "Apr", "May", "Jun", "Jul", "Aug", "Sept", "Oct", "Nov", "Desc"];
    // let minggu = ["week 1", "week 2", "week 3", "week 4"];
    let tanggal = time.getDate();
    let indexBulan = time.getMonth();
    let tahun = time.getFullYear();
    let hours = time.getHours();
    let minutes = time.getMinutes();

    if (hours < 10) {
        hours = "0" + hours
    }
    if (minutes < 10) {
        minutes = "0" + minutes
    }

    // if (tanggal <= 7) {
    //     minggu = minggu[0];
    //   } else if (tanggal <= 14) {
    //     minggu = minggu[1];
    //   } else if (tanggal <= 21) {
    //     minggu = minggu[2];
    //   } else if (tanggal <= 31) {
    //     minggu = minggu[3];
    //   }

    return `${tanggal} ${bulan[indexBulan]}  ${tahun} ${hours}:${minutes} WIB`;
    //   console.log(time);
}



