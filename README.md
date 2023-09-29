# fiap-tech-challenge-api

## Requirements

- Run swag to generate api documentation
```sh
swag init
```
- Run this command to generate dependency injection and mock test files
```sh
go generate ./...
```

- set env vars
```sh
export AUTH_SECRET={your_secret}
export DB_USER_PASS={db_password}
export DB_USER_NAME={db_name}
```
- some application configurations can be set into ``resources/config.yml``
- to build database (myqsl) container run ``docker-compose up -d``
---
### to run application
```sh
go run .
```

### to run tests
```sh
ginkgo -v ./...
```

### to access swagger doc
```
http://localhost:3000/swagger/index.html
```
