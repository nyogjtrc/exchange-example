# Exchange Example

a exchange api example

## Requirements

- Go 1.20

## Usage

command to execute API server
```
$ go run main.go
```

curl API example
```
$ curl 'localhost:8080/api/exchange?source=USD&target=JPY&amount=$1,525'                  1↵
{"amount":"$170,496.53","msg":"success"}%
```

## Development

testing command
```
$ go test ./...                                                                                                                                                                          130↵
?       github.com/nyogjtrc/exchange-example    [no test files]
ok      github.com/nyogjtrc/exchange-example/internal/rest      0.008s
```

coverage command
```
$ go test ./... -cover
?       github.com/nyogjtrc/exchange-example    [no test files]
ok      github.com/nyogjtrc/exchange-example/internal/rest      0.009s  coverage: 73.5% of statements
```
