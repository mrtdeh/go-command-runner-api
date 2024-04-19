# go-command-runner-api
Run command via REST API in server.
## Requirements
- go compiler v1.21

## Debug
first in once you must install dependencies used in this project by:
```sh
go mod tidy
```
then, you can simple run API server by:
```sh
go run main.go
```
you can test API with test.sh script:
```sh
./test.sh
```
## Build
build project in to bin directory:
```sh
go build -o bin
```
## Usage
```sh
./bin/exec-api -p 9090
...
```
request to server for running ***ls /*** command in server:
```sh
curl -XPOST http://localhost:9090/run -d '{"cmd":"ls /"}' -H "Content-Type: application/json" 
...
```
