#!/bin/bash

#Signin
#curl -d '{"username":"cfabrica46","password":"01234"}' -X POST http://localhost:8080/signin

curl -X GET http://localhost:8080/users
