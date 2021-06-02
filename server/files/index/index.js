var originBody = document.getElementById("body").innerHTML
let btnUserFriends = document.getElementById("btn-User-Friends")
var tokenString;

btnUserFriends.addEventListener("click", userFriends);

function userFriends() {

    const xhttpUserFriends = new XMLHttpRequest();

    let i = 0;

    xhttpUserFriends.onreadystatechange = function () {

        if (this.readyState === 4 && this.status === 200) {

            let title = document.getElementById("title")
            let form = document.getElementById("form")
            let btnIndex = document.getElementById("btn-index")
            let friends = JSON.parse(this.responseText)

            title.textContent = "List Of Friends"
            form.remove()
            btnIndex.textContent = "Back To Index"
            console.log(friends)

            let listFriends = document.createElement("ul")

            for (friend of friends) {
                const ulUser = document.createElement("ul")
                ulUser.innerHTML += `<li><h2>User: ${friend.Username}</h2></li>
                                    <li>ID: ${friend.ID}</li><br>
                                    <li>Password: ${friend.Password}</li><br>
                                    <li>Role: ${friend.Role}</li><br><br>`
                listFriends.appendChild(ulUser)
            }
            document.getElementById("body").appendChild(listFriends)

            btnIndex.addEventListener("click", () => backToIndex())

        } else if (this.status === 400) {

            btnUserFriends.remove();
            login();

        };

    };

    xhttpUserFriends.open("GET", "/api/v1/user/friends", true)
    xhttpUserFriends.setRequestHeader("Authorization", tokenString)
    xhttpUserFriends.send();
}

function login() {

    let body = document.getElementById("body");
    let title = document.getElementById("title");
    title.textContent = "Join Your Data";
    let form = document.createElement("form");

    form.id = "form"

    form.innerHTML += `Username: <input id="username" type="text" name="username"><br><br>
                        Password: <input id="password" type="password" name="password"><br><br>
                        <button id="btn-submit">Submit</button>`;

    body.appendChild(form);

    let btnSubmit = document.getElementById("btn-submit");

    btnSubmit.addEventListener("click", (e) => {
        e.preventDefault();

        const username = document.getElementById("username");
        const password = document.getElementById("password");

        let User = {
            ID: "",
            Username: username.value,
            Password: password.value,
            Role: "",
            Deadline: "",
            Token: "",
            Friends: []
        };

        const xhttpLogin = new XMLHttpRequest();

        xhttpLogin.onreadystatechange = function () {

            if (this.readyState === 4 && this.status === 200) {
                let token = JSON.parse(this.responseText);
                tokenString = token.Content;
                console.log(tokenString);
                userFriends();
            };

        };

        xhttpLogin.open("POST", "/api/v1/signin", true);
        xhttpLogin.send(JSON.stringify(User));
    })

}

function backToIndex() {

    let body = document.getElementById("body")
    body.innerHTML = originBody

    let btnUserFriends = document.getElementById("btn-User-Friends")
    btnUserFriends.addEventListener("click", () => userFriends());

}