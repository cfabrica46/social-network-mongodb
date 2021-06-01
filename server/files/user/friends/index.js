function traerDatos() {
    const xhttp = new XMLHttpRequest();

    xhttp.onreadystatechange = function () {
        if (this.readyState === 4 && this.status === 200) {
            let friends = JSON.parse(this.responseText)



        }
    }
    xhttp.open("GET", "/api/v1/user/friends", true);

    xhttp.send();
};

traerDatos();