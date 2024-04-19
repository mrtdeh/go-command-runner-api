# go-command-runner-api
Run command via REST API in the target node.

## Usage
```sh
./bin/exec-api -p 9090
...
```
and run commad :
```sh
curl -XPOST http://localhost:9090/run -d '{"cmd":"ls /"}' -H "Content-Type: application/json" 
...
```
