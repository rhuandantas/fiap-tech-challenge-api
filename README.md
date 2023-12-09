# fiap-tech-challenge-api

### Tech Challenge 2:
### Passos para homologação dos professores da Fiap

1. Instale a CLI do GCloud através das instruções do link a abaixo:
```
https://cloud.google.com/sdk/docs/install?hl=pt-br
```

2. Definar o projeto no GCP em que vai atuar:
```
gcloud config set project PROJECT_ID
```

3. Faça a autenticação e configure o docker:
```
sudo usermod -a -G docker ${USER}

gcloud auth login

gcloud auth configure-docker
```

Obs: Caso você execute o docker com o sudo, é necessário executar os comandos acima com o sudo também.

4. Execute os seguintes comandos abaixo para subir uma imagem da api no Container Registry do GCP:
```

docker build -t fiap-tech-challenge-api .

docker tag fiap-tech-challenge-api gcr.io/pos-tech-arquitetura/fiap-tech-challenge-api:latest

docker push gcr.io/pos-tech-arquitetura/fiap-tech-challenge-api:latest

```

Passo a passo em construção...



### Tech Challenge 1:
### Passos para homologação dos professores da Fiap

1. Faça o git clone do projeto:
```
git clone https://github.com/rhuandantas/fiap-tech-challenge-api.git
```

2. Cadastre a pasta /docker no File Sharing de seu Docker local.

3. Execute o seguinte comando na raiz do projeto:
```
docker-compose up --build
```

4. Importe as collections do Insomnia que estão no link abaixo:

https://github.com/rhuandantas/fiap-tech-challenge-api/blob/main/docs/insomnia_collection

Obs: Somente os status abaixo são válidos para executar os endpoints: atualiza status e lista por status:

1. "recebido"
2. "em_preparacao"
3. "pronto"
4. "finalizado"

---

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

### to update swagger files
```
swag init
```

### to update dependency injection file
```
delete wire_gen.go
go generate ./...
```
