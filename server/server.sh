#!/bin/bash

#Signin
#curl -X POST http://localhost:8080/signin -d '{"username":"cfabrica46","password":"01234"}'

#SignUp
#curl -X POST http://localhost:8080/signup -d '{"username":"cfabrica46","password":"789"}'

#Profile
#curl -X GET http://localhost:8080/user -H "Authorization-header: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZWFkLWxpbmUiOiJTYXQgTWF5IDI5IDEwOjE0OjI4IDIwMjEiLCJpZCI6IjYwYjFjNDkyNGFiMjkzZGU5NjFkYTBlNyIsInJvbGUiOiJhZG1pbiIsInVzZXJuYW1lIjoiY2ZhYnJpY2E0NiIsInV1aWQiOiI3OWVhOGVjNi0xMmNhLTRjZjktYmFjYi0zOTBkODI4Yzg0YjEifQ.82ZvIygphY_-rzdk_nEF48ZOgfhwhxWzIXTRNU8cnTU"

#ShowUsers
#curl -X GET http://localhost:8080/users

#Delete

#Update
#curl -X PUT http://localhost:8080/user -d '{"username":"uwu","password":"owo"}' -H "Authorization-header: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZWFkLWxpbmUiOiJTYXQgTWF5IDI5IDEwOjE0OjI4IDIwMjEiLCJpZCI6IjYwYjFjNDkyNGFiMjkzZGU5NjFkYTBlNyIsInJvbGUiOiJhZG1pbiIsInVzZXJuYW1lIjoiY2ZhYnJpY2E0NiIsInV1aWQiOiI3OWVhOGVjNi0xMmNhLTRjZjktYmFjYi0zOTBkODI4Yzg0YjEifQ.82ZvIygphY_-rzdk_nEF48ZOgfhwhxWzIXTRNU8cnTU"

#Show User's Posts
#curl -X GET http://localhost:8080/user/posts -H "Authorization-header: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZWFkLWxpbmUiOiJTYXQgTWF5IDI5IDEwOjE0OjI4IDIwMjEiLCJpZCI6IjYwYjFjNDkyNGFiMjkzZGU5NjFkYTBlNyIsInJvbGUiOiJhZG1pbiIsInVzZXJuYW1lIjoiY2ZhYnJpY2E0NiIsInV1aWQiOiI3OWVhOGVjNi0xMmNhLTRjZjktYmFjYi0zOTBkODI4Yzg0YjEifQ.82ZvIygphY_-rzdk_nEF48ZOgfhwhxWzIXTRNU8cnTU"

#Show User's Friends
curl -X GET http://localhost:8080/user/friends -H "Authorization-header: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZWFkLWxpbmUiOiJTYXQgTWF5IDI5IDEwOjE0OjI4IDIwMjEiLCJpZCI6IjYwYjFjNDkyNGFiMjkzZGU5NjFkYTBlNyIsInJvbGUiOiJhZG1pbiIsInVzZXJuYW1lIjoiY2ZhYnJpY2E0NiIsInV1aWQiOiI3OWVhOGVjNi0xMmNhLTRjZjktYmFjYi0zOTBkODI4Yzg0YjEifQ.82ZvIygphY_-rzdk_nEF48ZOgfhwhxWzIXTRNU8cnTU"