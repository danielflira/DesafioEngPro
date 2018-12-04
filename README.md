# Desafio EngPro

Olá! Esse é o desafio da Equipe DevOps da Engenharia Protheus. Esse repositório tem o diretório ./src com um código em Go, ele é uma aplicação web que da boas vindas para as pessoas ou para o mundo inteiro. Além da aplicação também existe o teste da função reponsável pelas boas vindas. Essa aplicação não possui nenhuma automatização apenas os fontes da aplicação, sua missão é criar os passos de automatização de acordo com o que está descrito nesse README.

### Criar Dockerfile

Crie um Dockerfile com os passos para testar a imagem então compilar o binário para deploy. Algumas dicas:

 - Coloque o diretório do src em /go/src/desafio
 - Existe a imagem golang
 - go test
 - go build
 - go install
 - Resultado da compilação esta em /bin/ com o nome de desafio

### Criar um compose

Crie um compose.yaml que inicie a nossa aplicação. Algumas dicas:

 - Aplicação exporta a porta 8080

### Criar Playbook

Crie o playbook para fazer o deploy da aplicação:

 - Enviar o compose.yaml para o servidor
 - Executar o docker stack deploy

### Criar Jenkinsfile

Crie um Jenkinsfile para automatizar esse processo, esse Jenkinsfile devera fazer os seguintes processos:

 - Executar o teste e build da imagem
 - Fazer o push da imagem
 - Fazer o deploy da imagem
