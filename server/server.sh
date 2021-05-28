#!/bin/bash

#Signin
#curl -d '{"username":"cfabrica46","password":"01234"}' -X POST http://localhost:8080/signin

#SignUp
#curl -d '{"username":"carlos","password":"789"}' -X POST http://localhost:8080/signup

#ShowUsers
curl -X GET http://localhost:8080/users

#Delete
#curl -H "Authorization-header: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZWFkLWxpbmUiOiJGcmkgTWF5IDI4IDAxOjQ0OjMyIDIwMjEiLCJpZCI6IjYwYjA4MzJkYzljMTI3M2Y2MjQwYTI5MiIsInJvbGUiOiJhZG1pbiIsInVzZXJuYW1lIjoiY2ZhYnJpY2E0NiIsInV1aWQiOiI5OGJhYzcxYi1lOGM4LTQ2NTEtOWI0Yy1jZDQ1MGI1M2U3MGEifQ.V6VCMxC-5R_Q0g2M9B8odB_Tsyd-03llaFBSrXyNL2s" -X DELETE http://localhost:8080/user