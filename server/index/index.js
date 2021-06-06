var originBody = document.body.innerHTML
var userUsername;

let btnUsers = document.getElementById("btn-Users")
let btnUserPosts = document.getElementById("btn-User-Posts")
let btnUserFriends = document.getElementById("btn-User-Friends")
let btnFriendPosts = document.getElementById("btn-Friend-Posts")
let btnFriendsPosts = document.getElementById("btn-Friends-Posts")
let btnSignIn = document.getElementById("btn-SignIn")
let btnLogOut = document.getElementById("btn-LogOut");
let btnIndex = document.getElementById("btn-index")

localStorage.removeItem("token")

btnUsers.addEventListener("click", users)
btnUserPosts.addEventListener("click", userPosts);
btnUserFriends.addEventListener("click", userFriends);
btnFriendPosts.addEventListener("click", friendPosts)
btnFriendsPosts.addEventListener("click", friendsPosts);
btnSignIn.addEventListener("click", login);
btnLogOut.addEventListener("click", logout)
btnIndex.addEventListener("click", backToIndex)

function users() {

    //    const xhttpUsers = new XMLHttpRequest();

    let nav = document.getElementById("nav")
    if (nav != null) {
        nav.remove()
    }

    let btnIndex = document.getElementById("btn-index")
    btnIndex.addEventListener("click", backToIndex)

    let title = document.getElementById("title")
    title.textContent = "All Users"

    fetch("/api/v1/users")
        .then(responsive => responsive.json())
        .then(users => {

            let main = document.getElementById("main")

            for (user of users) {
                const ulUser = document.createElement("ul")
                ulUser.classList.add("list-principal")
                ulUser.innerHTML += `<li><h2>User: ${user.User.Username}</h2></li>
                                              <li><h3>ID: ${user.User.ID}</h3></li>
                                              <li><h3>Role: ${user.User.Role}</h3></li>
                                              <li class="list-principal__post-label"><h3>Posts: </h3></li>`

                const ulPost = document.createElement("ul")
                ulPost.classList.add("list-principal__list-secundary")
                for (post of user.Posts) {
                    ulPost.innerHTML += `<ul class="list-principal__list-aux">
                                                          <li><h4>ID: ${post.ID}</h4></li>
                                                          <li>Content: ${post.Content}</li>
                                                          <li>Date: ${post.Date}</li>
                                                      </ul>`
                }
                ulUser.appendChild(ulPost)

                main.appendChild(ulUser)
            }
        })

    //    xhttpUsers.onreadystatechange = function () {
    //        if (this.readyState === 4 && this.status === 200) {
    //            let users = JSON.parse(this.responseText)
    //
    //            let main = document.getElementById("main")
    //
    //            for (user of users) {
    //                const ulUser = document.createElement("ul")
    //                ulUser.classList.add("list-principal")
    //                ulUser.innerHTML += `<li><h2>User: ${user.User.Username}</h2></li>
    //                                    <li><h3>ID: ${user.User.ID}</h3></li>
    //                                    <li><h3>Role: ${user.User.Role}</h3></li>
    //                                    <li class="list-principal__post-label"><h3>Posts: </h3></li>`
    //
    //                const ulPost = document.createElement("ul")
    //                ulPost.classList.add("list-principal__list-secundary")
    //                for (post of user.Posts) {
    //                    ulPost.innerHTML += `<ul class="list-principal__list-aux">
    //                                            <li><h4>ID: ${post.ID}</h4></li>
    //                                            <li>Content: ${post.Content}</li>
    //                                            <li>Date: ${post.Date}</li>
    //                                        </ul>`
    //                }
    //                ulUser.appendChild(ulPost)
    //
    //                main.appendChild(ulUser)
    //            }
    //
    //        }
    //    }
    //    xhttpUsers.open("GET", "/api/v1/users", true);
    //    xhttpUsers.send();
}

function userPosts() {
    const xhttpUserPosts = new XMLHttpRequest();

    let nav = document.getElementById("nav")
    if (nav != null) {
        nav.remove()
    }

    let btnIndex = document.getElementById("btn-index")
    btnIndex.addEventListener("click", backToIndex)

    xhttpUserPosts.onreadystatechange = function () {

        if (this.readyState === 4 && this.status === 200) {

            let title = document.getElementById("title")
            let posts = JSON.parse(this.responseText)

            title.textContent = "Your Posts"

            let main = document.getElementById("main")

            let listPost = document.createElement("ul")

            for (post of posts) {
                const ulUser = document.createElement("ul")
                ulUser.classList.add("list-principal")

                ulUser.innerHTML += `<li><h3>ID: ${post.ID}</h3></li>
                                     <li>Content: ${post.Content}</li>
                                     <li>Date: ${post.Date}</li>`

                listPost.appendChild(ulUser)
            }
            main.appendChild(listPost)


        } else if (this.status === 400) {

            document.getElementById("title").textContent = "Error: Necesita autenticarse"

        };
    };

    xhttpUserPosts.open("GET", "/api/v1/user/posts", true)
    xhttpUserPosts.setRequestHeader("Authorization", localStorage.getItem("token"))
    xhttpUserPosts.send();
}

function userFriends() {
    const xhttpUserFriends = new XMLHttpRequest();

    let nav = document.getElementById("nav")
    if (nav != null) {
        nav.remove()
    }

    let btnIndex = document.getElementById("btn-index")
    btnIndex.addEventListener("click", backToIndex)

    xhttpUserFriends.onreadystatechange = function () {

        if (this.readyState === 4 && this.status === 200) {

            let title = document.getElementById("title")
            let friends = JSON.parse(this.responseText)

            title.textContent = "Your Friends"

            let main = document.getElementById("main")

            let listFriends = document.createElement("ul")

            for (friend of friends) {
                const ulUser = document.createElement("ul")
                ulUser.classList.add("list-principal")
                ulUser.innerHTML += `<li><h2>User: ${friend.Username}</h2></li>
                                     <li>ID: ${friend.ID}</li>
                                     <li>Role: ${friend.Role}</li>`
                listFriends.appendChild(ulUser)
            }
            main.appendChild(listFriends)


        } else if (this.status === 400) {

            document.getElementById("title").textContent = "Error: Necesita autenticarse"

        };
    };

    xhttpUserFriends.open("GET", "/api/v1/user/friends", true)
    xhttpUserFriends.setRequestHeader("Authorization", localStorage.getItem("token"))
    xhttpUserFriends.send();
}

function friendPosts() {

    let body = document.body;
    document.getElementById("title").textContent = "Enter your friend's username";
    let nav = document.getElementById("nav")
    if (nav != null) {
        nav.remove()
    }

    let btnIndex = document.getElementById("btn-index")
    btnIndex.addEventListener("click", backToIndex)

    if (localStorage.getItem("token") === null) {
        document.getElementById("title").textContent = "Error: Necesita autenticarse"
        return
    }

    let form = document.createElement("form");

    form.id = "form"
    form.classList.add("form")

    form.innerHTML += `Friend's Username: <input id="friendUsername" class="form__input-text" type="text" name="friendUsername"><br><br>
                        <button id="btn-submit" class="form__submit">Submit</button>`;

    body.appendChild(form);

    let btnSubmit = document.getElementById("btn-submit");
    btnSubmit.addEventListener("click", (e) => {
        e.preventDefault();

        const friendUsername = document.getElementById("friendUsername").value;

        const xhttpFriendPosts = new XMLHttpRequest();

        xhttpFriendPosts.onreadystatechange = function () {

            let main = document.getElementById("main")


            if (this.readyState === 4 && this.status === 200) {
                let friendPosts = JSON.parse(this.responseText);

                let title = document.getElementById("title")
                title.textContent = `${friendUsername}'s Posts`
                for (friendPost of friendPosts.Posts) {
                    const ulFriendPosts = document.createElement("ul")
                    ulFriendPosts.classList.add("list-principal")
                    ulFriendPosts.innerHTML += `<li><h3>ID: ${friendPost.ID}</h3></li>
                                    <li>Content: ${friendPost.Content}</li>
                                    <li>Date: ${friendPost.Date}</li>`
                    main.appendChild(ulFriendPosts)

                }


            } else {
                form.remove()
                title.textContent = "Error: The Username Entered Does Not Belong To One Of Your Friends"
            }

        };
        xhttpFriendPosts.open("GET", `/api/v1/friend/${friendUsername}/posts`, true);
        xhttpFriendPosts.setRequestHeader("Authorization", localStorage.getItem("token"))
        xhttpFriendPosts.send();
    })


}

function friendsPosts() {

    const xhttpFriendsPosts = new XMLHttpRequest();

    let nav = document.getElementById("nav")
    if (nav != null) {
        nav.remove()
    }

    let btnIndex = document.getElementById("btn-index")
    btnIndex.addEventListener("click", backToIndex)

    let title = document.getElementById("title")
    title.textContent = "Your Friends' Posts"

    xhttpFriendsPosts.onreadystatechange = function () {
        if (this.readyState === 4 && this.status === 200) {
            let friendsPosts = JSON.parse(this.responseText)

            let main = document.getElementById("main")


            for (friendPost of friendsPosts) {
                const ulfriendPost = document.createElement("ul")
                ulfriendPost.classList.add("list-principal")
                ulfriendPost.innerHTML += `<li><h2>Author: ${friendPost.Author}</h2></li>
                                    <li class="list-principal__post-label"><h3>Posts: </h3></li>
                                    <ul class="list-principal__list-aux">
                                        <li><h3>ID: ${friendPost.Post.ID}</h3></li>
                                        <li>Content: ${friendPost.Post.Content}</li>
                                        <li>Date: ${friendPost.Post.Date}</li>
                                    </ul>`
                main.appendChild(ulfriendPost)

            }

        } else if (this.status === 400) {

            document.getElementById("title").textContent = "Error: Necesita autenticarse"

        };
    }
    xhttpFriendsPosts.open("GET", "/api/v1/friends/posts", true);
    xhttpFriendsPosts.setRequestHeader("Authorization", localStorage.getItem("token"))
    xhttpFriendsPosts.send();

}

function login() {

    let body = document.body;
    let title = document.getElementById("title");
    title.textContent = "Join Your Data";
    let form = document.createElement("form");

    let btnIndex = document.getElementById("btn-index")
    btnIndex.addEventListener("click", backToIndex)

    let nav = document.getElementById("nav")
    if (nav != null) {
        nav.remove()
    }

    form.id = "form"
    form.classList.add("form")

    form.innerHTML += `Username: <input id="username" class="form__input-text" type="text" name="username"><br><br>
                    Password: <input id="password" class="form__input-text" type="password" name="password"><br><br>
                        <button id="btn-submit" class="form__submit">Submit</button>`;

    body.appendChild(form);

    let btnSubmit = document.getElementById("btn-submit");

    btnSubmit.addEventListener("click", (e) => {
        e.preventDefault();

        const username = document.getElementById("username");
        const password = document.getElementById("password");

        let User = {
            Username: username.value,
            Password: password.value
        };

        const xhttpLogin = new XMLHttpRequest();

        xhttpLogin.onreadystatechange = function () {

            if (this.readyState === 4 && this.status === 200) {
                let token = JSON.parse(this.responseText);
                userUsername = username.value;
                localStorage.setItem("token", token.Content);
                backToIndex();
            } else {
                form.remove()
                title.textContent = "Error: Username and/or Password Invalid"
            }

        };

        xhttpLogin.open("POST", "/api/v1/signin", true);
        xhttpLogin.send(JSON.stringify(User));
    })

}

function logout() {
    const xhttpLogout = new XMLHttpRequest();
    xhttpLogout.onreadystatechange = function () {

        if (this.readyState === 4 && this.status === 200) {
            userUsername = ""
            localStorage.removeItem("token")
            backToIndex();
        } else {
            backToIndex();
        }

    }

    xhttpLogout.open("GET", "/api/v1/logout", true);
    xhttpLogout.setRequestHeader("Authorization", localStorage.getItem("token"))
    xhttpLogout.send();
}

function backToIndex() {

    let body = document.body;
    body.innerHTML = originBody

    if (userUsername == undefined || userUsername == "") {
        document.getElementById("title").textContent = "Index"
    } else {
        document.getElementById("title").textContent = `${userUsername}'s Profile`
    }

    let btnUserFriends = document.getElementById("btn-User-Friends");
    let btnUserPosts = document.getElementById("btn-User-Posts");
    let btnFriendsPosts = document.getElementById("btn-Friends-Posts");
    let btnUsers = document.getElementById("btn-Users");
    let btnFriendPosts = document.getElementById("btn-Friend-Posts");
    let btnSignIn = document.getElementById("btn-SignIn");
    let btnLogOut = document.getElementById("btn-LogOut");
    let btnIndex = document.getElementById("btn-index");

    btnUserFriends.addEventListener("click", userFriends);
    btnUserPosts.addEventListener("click", userPosts);
    btnFriendsPosts.addEventListener("click", friendsPosts);
    btnUsers.addEventListener("click", users)
    btnFriendPosts.addEventListener("click", friendPosts)
    btnSignIn.addEventListener("click", login);
    btnLogOut.addEventListener("click", logout)
    btnIndex.addEventListener("click", backToIndex)

}