# fiap-tech-challenge-api

### Tech Challenge 2:
### Passos para homologação dos professores da Fiap

Foi utilizado o Ubuntu como sistema operacional no passo a passo abaixo, além disso, o computador estava com o docker instalado:

1. Crie uma conta e um projeto na plataforma do Google Cloud Platform (GCP), como exemplo neste passo a passo vamos supor que o nome do projeto criado é "fiap-pos-tech-arquitetura";

2. Ative o Container Registry no GCP;

3. Instale a CLI do GCloud através das instruções do link a abaixo:
```
https://cloud.google.com/sdk/docs/install?hl=pt-br
```

4. Definar o projeto no GCP em que vai atuar:
```
# gcloud config set project PROJECT_ID
gcloud config set project fiap-pos-tech-arquitetura
```

5. Faça a autenticação e configure o docker:
```
sudo usermod -a -G docker ${USER}

gcloud auth login

gcloud auth configure-docker
```

Obs: Caso você execute o docker com o sudo, é necessário executar os comandos acima com o sudo também.

6. Execute os seguintes comandos abaixo na raiz do projeto para subir uma imagem da api no Container Registry do GCP, considere que no exemplo o nome da imagem será "fiap-tech-challenge-api":
```

docker build -t fiap-tech-challenge-api .

docker tag fiap-tech-challenge-api gcr.io/fiap-pos-tech-arquitetura/fiap-tech-challenge-api:latest

docker push gcr.io/fiap-pos-tech-arquitetura/fiap-tech-challenge-api:latest

```

7. Após isso, no GCP ative o Kubernetes Engine API, este passo poderá levar alguns minutos;

8. Após isso vá até a tela de Clusters no menu "Kubernetes Engine->Clusters" e clique em "CREATE", caso a tela de criação de cluster autopilot seja aberta, clique no botão superior direiro com o label "switch to standard cluster", é necessário que seja criado um cluster standard pois o servidor de métricas usado pelo HPA não funcionará em um cluster autopilot. Neste passo a passo, um cluster com as configurações padrões já serão suficiente. Após o procedimento de criação do cluster ser finalizado por você, o GCP poderá levar alguns minutos para finalizar;

9. Após o GCP terminar de criar o cluster, clique sobre o nome do mesmo, e na tela que se abrir clique no botão no topo da página chamado "CONNECT", um modal irá se abrir, para facilitar a conexão com o cluster e não termos problemas na instalação de algumas bibliotecas na máquina local, clique em "RUN IN CLOUD SHELL";

10. Será aberto no rodapé uma janela que se assemelha a um terminal, após o comando de conexão com o cluster aparecer automaticamente, apenas clique em "ENTER", uma janela de confirmação com o título "Authorize Cloud Shell" irá aparecer, apenas confirme clicando em "AUTHORIZE";

11. Voltando ao projeto, altere o nome da imagem no arquivo infra/api-go/api-deployment.yaml para o nome da imagem que você subiu no Container Registry

11. Crie um arquivo .zip da pasta chamada "infra" deste projeto, o nome do arquivo deverá ficar como "infra.zip";

12. De volta ao terminal do Cloud Shell do GCP, clique no ícone de três pontos no canto superior direito do mesmo e depois na opção "UPLOAD", faça o upload do arquivo infra.zip que acabou de criar;

13. Ainda no Cloud Shell, execute a seguinte sequência de comandos abaixo:
```
unzip infra.zip

cd infra/

kubectl apply -f secrets.yaml

cd metrics-server/

# Metrics server downloaded from https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
kubectl apply -f components.yaml

cd ../mysql/

kubectl apply -f mysql-pv.yaml

kubectl apply -f mysql-svc.yaml

kubectl apply -f mysql-deployment.yaml

cd ../api-go/

kubectl apply -f api-hpa.yaml

kubectl apply -f api-svc.yaml

kubectl apply -f api-deployment.yaml
```

14. O servidor de métricas pode levar de 5 minutos ou até mais para começar a coletar as métricas corretamente, para verificar se o mesmo já está funcionando execute o comando abaixo:
```
kubectl get hpa --watch
```
Caso a coluna "TARGETS" esteja com um valor de "<unknown>/80%", então o mesmo ainda não está funcionando, quando começar a aparecer algum número antes do /80%, por exemplo assim "0%/80%", então o mesmo começou a funcionar;

15. De volta a interface do GCP, vá até o menu "Kubernetes Engine->Gateways, Services & Ingress" e copie o valor da coluna "Endpoints" da linha com o nome "api-svc", o valor copiado será algo semelhante a "http://34.41.255.86:3000/";


A verficação de um endpoint com este valor seria assim por exemplo: http://34.41.255.86:3000/liveness


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
