#!/bin/bash

#Signin
#curl -d '{"username":"cfabrica46","password":"01234"}' -X POST http://localhost:8080/signin

#SignUp
#curl -d '{"username":"cfabrica46","password":"789"}' -X POST http://localhost:8080/signup

#Profile
#curl -H "Authorization-header: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZWFkLWxpbmUiOiJGcmkgTWF5IDI4IDE5OjUwOjM0IDIwMjEiLCJpZCI6IjYwYjE4MWJiNjNjMGE0ZjUyMGNiNTYyYiIsInJvbGUiOiJhZG1pbiIsInVzZXJuYW1lIjoiY2ZhYnJpY2E0NiIsInV1aWQiOiI0YWJiN2M2Mi0yNGU2LTRkOGEtYWQ5My1jZjJkYWE5NmI4N2IifQ.0CIG2m9PVsK1gpAG3cxorQohunPObAl9EiEV811HHnA" -X GET http://localhost:8080/user

#ShowUsers
#curl -X GET http://localhost:8080/users

#Delete
#curl -H "Authorization-header: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZWFkLWxpbmUiOiJGcmkgTWF5IDI4IDE5OjMwOjM5IDIwMjEiLCJpZCI6IjYwYjE3Y2IyZWQwMjhjYjNiZmY0MTQ2NiIsInJvbGUiOiJhZG1pbiIsInVzZXJuYW1lIjoiY2ZhYnJpY2E0NiIsInV1aWQiOiI5NGI2NzYzZi05NjY3LTRjMjQtYTkzNS1kM2E3YjBmNzUwMDMifQ.dZXdgWnGSyHg7M1FvLffDNwKHIc16JjzgVOW04GLEiY" -X DELETE http://localhost:8080/user