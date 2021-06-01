console.log("correcto");

function traerDatos() {

    console.log("dentro de la funcion")

    const xhttp = new XMLHttpRequest();

    xhttp.open("GET", "./user.json", true);

    xhttp.send();

    xhttp.onreadystatechange = function () {
        if (this.readyState === 4 && this.status === 200) {
            console.log(this.responseText)
        }
    }
};

traerDatos();

