## Instructions

In order to run the serivice without compiling
```
go run api.go
```
In oder to run the service compiling the sourcode before
```
go build api.go
./api
```
Once the service is running use 
```
go test
```
inside the main folder to run the end to end tests

use the same command inside the lib folter
to run the unittests


## Disclaimer

This is a skill showcase of my golang only skill:
 - I have not included odm ( I generally use mgo for MongoDB)
 - I have used global structures instead that will vanish at the end at the end of the
   of the instance that is running this code
 - To handle the api calls I have used the library fasthttp