# Teste Gobrax
Repositório com meu projeto para desafio da empresa Gobrax, link para descrição do desafio: https://gitlab.com/gobrax-public/backend-challenges1

## Modo de construção
Os pré requisistos para construção do projeto são __Docker__ e __Docker Compose__ instalado na máquina, sugiro a utilização do sistema operacional Linux ou WSL(no Windows). O primeiro passo a seguir é a construição do .env, tem um arquivo de exemplo chamado .env.example, nele contém exemplos das configurações que o sistema precisa para funcionar, se o texto que existe dentro dele for copiado para um arquivo na mesma pasta com o nome .env já é o suficiente para que tudo esteja certo.
Após a configuração do .env, com o comando a seguir o docker irá baixar e construir as imagens de MySQL e do sistema e começar a rodar os containers:


    sudo docker compose up

O entrypoint ouvirá na porta que a env **LISTENER_PORT** for colocado, por padrão no exemplo é a 8080.

Para subir somente o MySQL ou o Entrpoint deve-se colocar o nome do serviço (mysql ou entrypoint) após o comando para que somente aquela parte do projeto seja construida:

    sudo docker compose up entrypoint

## Endpoints criados
Existe um arquivo com uma colection do Postman que pode ser importada com o nome de Gobrax.postman_collection.json na pasta root do repositório, nela contém todos os endpoints já incluidos e deve somente ser criado uma variável com o host. Ex: http://localhost:8080

Para CRUD do motorista foram criados os endpoints:
- 	**GET** /drivers - Listagem de todos os motoristas paginado (aceita campos query _page_ e _per_page_)
-	**POST** /driver - Cadastro de um novo motorista, deve receber um corpo no formato json Ex: {"name": "Jose Bezerra","cnh": "12345678901"}
-	**GET** /driver/:id - Pegar informações de um motorista, aceita variável com o identificador no caminho
-	**PUT** /driver/:id - Atualização de um motorista, aceita variável com o identificador no caminho e deve conter um corpo com todos os campos de um cadastro
-	**DELETE** /driver/:id - Delete de um motorista

Para CRUD do caminhão foram criados os endpoints:
- 	**GET** /trucks - Listagem de todos os caminhão paginado (aceita campos query _page_ e _per_page_)
-	**POST** /truck - Cadastro de um novo caminhão, deve receber um corpo no formato json Ex: {"brand": "VOLVO","plate": "BBB1234"}
-	**GET** /truck/:id - Pegar informações de um caminhão, aceita variável com o identificador no caminho
-	**PUT** /truck/:id - Atualização de um caminhão, aceita variável com o identificador no caminho e deve conter um corpo com todos os campos de um cadastro
-	**DELETE** /truck/:id - Delete de um caminhão

Para CRUD da relação caminhão com motorista foram criados os endpoints:
- 	**GET** /trucks - Listagem de todos as relações paginado (aceita campos query _page_ e _per_page_)
-	**POST** /truck - Cadastro de uma nova relação, deve receber um corpo no formato json com identificadores válidos Ex: {"driver_id": "3c993bd1-34c9-4cd6-a2b8-6b122a2f5767","truck_id": "02027a95-32af-48f9-a03d-ec29426a3ed9"}
-	**GET** /truck/:id - Pegar informações de uma relação, aceita variável com o identificador no caminho
-	**DELETE** /truck/:id - Delete de uma relação
