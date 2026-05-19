# simple-go-server

Tiny HTTP server in Go that returns a JSON payload on `/`.

## Run

```sh
go run main.go
```

Then:

```sh
curl http://localhost:4444/
# {"Name":"Hello","Description":"World","Url":"localhost:4444"}
```

Listens on `:4444`.
