#### mylibrary


##### Start the server

```go
go run main.go
```

##### Add Books

```
curl -X POST -H "Content-Type: application/json" -d '{"title":"Alchemist","author":"Paulo"}' http://localhost:5000/addbooks
```


##### Get All Books

```
curl http://localhost:5000/allbooks | jq .
```
