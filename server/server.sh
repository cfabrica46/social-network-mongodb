#!/bin/bash

#Signin
#curl -X POST http://localhost:8080/signin -d '{"username":"cfabrica46","password":"01234"}'

#SignUp
#curl -X POST http://localhost:8080/signup -d '{"username":"cfabrica46","password":"789"}'

#Profile
#curl -X GET http://localhost:8080/user -H "Authorization-header: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZWFkLWxpbmUiOiJGcmkgTWF5IDI4IDE5OjUwOjM0IDIwMjEiLCJpZCI6IjYwYjE4MWJiNjNjMGE0ZjUyMGNiNTYyYiIsInJvbGUiOiJhZG1pbiIsInVzZXJuYW1lIjoiY2ZhYnJpY2E0NiIsInV1aWQiOiI0YWJiN2M2Mi0yNGU2LTRkOGEtYWQ5My1jZjJkYWE5NmI4N2IifQ.0CIG2m9PVsK1gpAG3cxorQohunPObAl9EiEV811HHnA"

#ShowUsers
#curl -X GET http://localhost:8080/users

#Delete

#Update
#curl -X PUT http://localhost:8080/user -d '{"username":"uwu","password":"owo"}' -H "Authorization-header: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZWFkLWxpbmUiOiJGcmkgTWF5IDI4IDIxOjU4OjA4IDIwMjEiLCJpZCI6IjYwYjE5ZjllYmZhYjJlNzI4MTJkOTI4MiIsInJvbGUiOiJhZG1pbiIsInVzZXJuYW1lIjoiY2ZhYnJpY2E0NiIsInV1aWQiOiI4ODhlNzJhNC0yZTI4LTQ5OWMtOWYwMC1hNGIwNDgxMWI3MDUifQ.LDG43hC3iUExZ5LNTfnwIWa2WYE6xqmlRgdro9vtu70"

#Show User's Posts
#curl -X GET http://localhost:8080/user/posts -H "Authorization-header: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZWFkLWxpbmUiOiJTYXQgTWF5IDI5IDAwOjU3OjU0IDIwMjEiLCJpZCI6IjYwYjFjNDkyNGFiMjkzZGU5NjFkYTBlNyIsInJvbGUiOiJhZG1pbiIsInVzZXJuYW1lIjoiY2ZhYnJpY2E0NiIsInV1aWQiOiI0MDBkOGJlYi0wMWYxLTQwYzYtOTk3ZS05NDFhYTZhZmY2ZGQifQ.-BaUr5QYVBrUoqXCWNiEtOz5bVLhdkdsVnnVbmQk3Jw"