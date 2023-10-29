# fiap-tech-challenge-api

## Requirements

### install libs
```sh
go install github.com/google/wire/cmd/wire@latest
go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/swag
go get -u github.com/google/wire/cmd/wire
```

- Run swag to generate api documentation
```sh
swag init
```
- Run this command to generate dependency
```sh
go generate ./... or wire ./...
```
---
### to run application
```sh
go run .
```

### to access swagger doc
```
http://localhost:3000/docs/index.html
```
