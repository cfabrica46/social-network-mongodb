const btn = document.getElementById("btn")
const username = document.getElementById("username")
const password = document.getElementById("password")

btn.addEventListener("click", (e) => {
    e.preventDefault();

    let User = {
        ID: "",
        Username: username.value,
        Password: password.value,
        Role: "",
        Deadline: "",
        Token: "",
        Friends: []
    };

    const xhttp = new XMLHttpRequest();

    xhttp.onreadystatechange = function () {

        if (this.readyState === 4 && this.status === 200) {
            let token = JSON.parse(this.responseText)
            console.log(token)

            const xhttp2 = new XMLHttpRequest;

            xhttp2.onreadystatechange = function () {
                if (this.readyState === 4 && this.status === 200) {
                    const title = document.getElementById("title")
                    const form = document.getElementById("form")
                    let friends = JSON.parse(this.responseText)

                    title.textContent = "List Of Friends"
                    form.remove()
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
                }
            }

            xhttp2.open("GET", "/api/v1/user/friends", true)
            xhttp2.setRequestHeader("Authorization", token.Content)
            xhttp2.send();
        }
    }

    xhttp.open("POST", "/api/v1/signin", true);
    xhttp.send(JSON.stringify(User));

})

