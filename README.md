# fiap-tech-challenge-api

### this go application needs go version 1.20 or later

---

### export env variables
```
export DB_HOST=
export DB_PORT=
export DB_NAME=
export DB_PASS=
export DB_USER=
```

### to run application locally
```sh
go mod download
go run .
```

### to run via docker
```sh
docker compose up
```

### to access swagger doc
```
http://localhost:3000/docs/index.html
```

## Development
### Requirements

### install libs
```sh
go install github.com/google/wire/cmd/wire@latest
go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/swag
go get -u github.com/google/wire/cmd/wire
```