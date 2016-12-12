#Simple post api in Go
This is just a simple api receiving POST requests that get inserted into a
MySQL database in Go, using Go's web framework **Gin**.
Requests are handled asynchronously on independent goroutines.


example request:

`curl -X POST http://127.0.0.1:8080/post --header "Content-Type:application/json" -d'{"name": "adrian", "time": "2016-01-01 00:01:01"}'`
