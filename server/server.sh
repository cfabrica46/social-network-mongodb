#!/bin/bash

#Signin
#curl -d '{"username":"cfabrica46","password":"01234"}' -X POST http://localhost:8080/signin

#SignUp
#curl -d '{"username":"carlos","password":"789"}' -X POST http://localhost:8080/signup

#Profile
#curl -H "Authorization-header: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZWFkLWxpbmUiOiJGcmkgTWF5IDI4IDE5OjMxOjM2IDIwMjEiLCJpZCI6IjYwYjE3ZDU1N2JjM2VhYWY1ZGRjYWVjZCIsInJvbGUiOiJhZG1pbiIsInVzZXJuYW1lIjoiY2ZhYnJpY2E0NiIsInV1aWQiOiI4NmU3NzdhMC05ODVkLTQ1NTctODY4YS1mNWUxZmI4MjA2ZTQifQ.YCKbbWYhd0JgEdZESluafBFj2odL4r4ozgmCApxeWyw" -X GET http://localhost:8080/user

#ShowUsers
#curl -X GET http://localhost:8080/users

#Delete
#curl -H "Authorization-header: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZWFkLWxpbmUiOiJGcmkgTWF5IDI4IDE5OjMwOjM5IDIwMjEiLCJpZCI6IjYwYjE3Y2IyZWQwMjhjYjNiZmY0MTQ2NiIsInJvbGUiOiJhZG1pbiIsInVzZXJuYW1lIjoiY2ZhYnJpY2E0NiIsInV1aWQiOiI5NGI2NzYzZi05NjY3LTRjMjQtYTkzNS1kM2E3YjBmNzUwMDMifQ.dZXdgWnGSyHg7M1FvLffDNwKHIc16JjzgVOW04GLEiY" -X DELETE http://localhost:8080/user