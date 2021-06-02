const btn = document.getElementById("btn")
const username = document.getElementById("username")
const password = document.getElementById("password")

btn.addEventListener("click", (e) => {
    e.preventDefault();

    let User = {
        //        ID: null,
        Username: username.value,
        Password: password.value,
        //        Role: null,
        //        Deadline: null,
        //        Token: null,
        //        Friends: null
    };

    console.log(User);

    const xhttp = new XMLHttpRequest();

    xhttp.onreadystatechange = () => {
        if (this.readyState === 4 && this.status === 200) {
            console.log(this.responseText);
        }
    }

    xhttp.open("POST", "/api/v1/signin", true);
    xhttp.send(JSON.stringify(User));

})