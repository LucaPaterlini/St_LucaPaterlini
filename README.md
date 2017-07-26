## Disclaimer

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

Due to to specific goal of this task I've kept all the informations
inside local structures instead of using a database.
I've assumed that a list of users and items has been previously creted
with a unique id for each user and an unique id for
each item.
I've assumed that the access controll was not part of this task.
I've stored the value using a integer an representing the value in
pennies to not create errors due to the nature of the representation
of the floating point.
I've developed this task using echo framework and gogland as IDE


To run the tests I've used go test -parallel 1

I'm really aware that in reality all of them can occur in any order
but I cannot test all the race conditions that are in the nature
of the bids
