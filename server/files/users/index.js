function traerDatos() {
    const xhttp = new XMLHttpRequest();

    xhttp.onreadystatechange = function () {
        if (this.readyState === 4 && this.status === 200) {
            let users = JSON.parse(this.responseText)

            for (user of users) {
                const ulUser = document.createElement("ul")
                ulUser.innerHTML += `<li><h2>User: ${user.User.Username}</h2></li>
                                    <li>ID: ${user.User.ID}</li><br>
                                    <li>Password: ${user.User.Password}</li><br>
                                    <li>Role: ${user.User.Role}</li>
                                    <li><h3>Posts: </h3></li>`

                const ulPost = document.createElement("ul")
                for (post of user.Posts) {
                    console.log(post)
                    ulPost.innerHTML += `<ul>
                                            <li><h4>ID: ${post.ID}</h4></li>
                                            <li>Content: ${post.Content}</li><br><br>
                                        </ul>`
                }
                ulUser.appendChild(ulPost)

                let listUser = document.getElementById("list-user")
                listUser.appendChild(ulUser)
            }

        }
    }
    xhttp.open("GET", "/api/v1/users", true);

    xhttp.send();
};

traerDatos();