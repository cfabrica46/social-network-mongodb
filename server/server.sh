#!/bin/bash

#Signin
curl -d '{"username":"cfabrica46","password":"01234"}' -X POST http://localhost:8080/signin

#SignUp
#curl -d '{"username":"carlos","password":"789"}' -X POST http://localhost:8080/signup

#ShowUsers
#curl -X GET http://localhost:8080/users

#Delete
#curl -H "Authorization-header: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZWFkLWxpbmUiOiJGcmkgTWF5IDI4IDAwOjE1OjIxIDIwMjEiLCJpZCI6IjYwYjA2ZTRlMWRiMzk2YWQ4MDhhOWM5ZiIsInJvbGUiOiJhZG1pbiIsInVzZXJuYW1lIjoiY2ZhYnJpY2E0NiIsInV1aWQiOiIxNDIyNTVmMC1lM2ViLTQ5ZDItOTQ5YS1lNGRlNDZmM2YxYWUifQ.5-OuyYm20Fe9dxqkjB8eLZe7x6dwLfaWUnDTEEGALzk" -X DELETE http://localhost:8080/user