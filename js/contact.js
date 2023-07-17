// alert("isi biar kita bisa call ing an")

// function dan variabel dari data form contact
function submitData() {
    let name = document.getElementById("input-name").value;
    let email = document.getElementById("input-email").value;
    let phone = document.getElementById("input-phone").value;
    let subject = document.getElementById("input-subject").value;
    let message = document.getElementById("input-message").value;


    // condition data
    if (name == "") {
        return alert("Nama Harus Di isii");
    } else if (email == "") {
        return alert("Email Harus Di isii");
    } else if (phone == "") {
        return alert("Phone Harus Di isii");
    } else if (subject == "") {
        return alert("Subject Harus Di Pilih");
    } else if (message == "") {
        return alert("Message Harus Di ISII");
    };

    document.getElementById("input-name").value = ""
    document.getElementById("input-email").value = ""
    document.getElementById("input-phone").value = ""
    document.getElementById("input-subject").value = ""
    document.getElementById("input-message").value = ""


    console.log(name);
    console.log(email);
    console.log(phone);
    console.log(subject);
    console.log(message);

    let emailReceiver = "andigagaga1@gmail.com"


    let a = document.createElement("a");
    a.href = `mailto:${emailReceiver}?subject= ${subject} &body=Halo, Perkenalkan Nama Saya ${name}, ${message};    Silahkan Hubungi Saya Di Nomor ${phone}, Terima Kasih.`;
    a.click();

    // untuk mengisi data secara object keknyaa
    let objecter = {
        name,
        email,
        phone,
        subject,
        message,
    };
    console.log(objecter)

}






