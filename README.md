# API REST para gerenciamento de cadastro de alunos

Este é um projeto de exemplo em Go para construir uma API REST destinada à manipulação de cadastro de alunos. A API permite que você liste, crie, atualize e delete alunos usando os métodos HTTP GET, POST, PUT e DELETE.

A implementação utiliza alguns packages úteis, tais como:
- [Gin Web Framework](https://github.com/gin-gonic/gin): Framework web rápido e flexível para Go;
- [Gorm](https://gorm.io/index.html): Biblioteca ORM para Go que disponibiliza drivers para diversos bancos de dados (neste projeto, utilizado o para `PostgreSQL`);
- [Validator.v2](https://github.com/go-playground/validator): Biblioteca para validação de structs em Go.


# Iniciando

## Configuração
- Clonar repositório:
```git clone https://github.com/sirio-neto/gin-rest-api.git```
- Acessar diretório projeto
```cd gin-rest-api```
- Criar arquivo `.env` com base em arquivo de exemplo, para configuração de variáveis de ambiente:
```cp .env_example .env```
- Iniciar containers de serviços (`postgres` e `pgadmin4`) docker para criação e administração do banco de dados PostgreSQL:
```docker compose up -d```

## Utilização
- O projeto pode ser iniciado executando:
```go run main.go```
- Após inicialização, a conexão com banco de dados será estabelecida e as migrations de cada model serão executadas através do `gorm`.
- A API REST ficará acessível em http://localhost:8002, com seguintes endpoints disponíveis:
	- `GET /:nome`: Retorna uma mensagem de boas-vindas com o nome fornecido.
	- `GET /students`: Lista todos os alunos cadastrados.
	- `GET /students/cpf/:cpf`: Retorna o aluno com o CPF informado.
	- `GET /students/:id`: Retorna o aluno com o ID informado.
	- `DELETE /students/:id`: Remove o aluno com o ID informado.
	- `PATCH /students/:id`: Atualiza os dados do aluno com o ID informado.
	- `POST /students`: Cria um novo aluno.

## Testes
É possível executar os testes automatizados incluídos no projeto. Para isso:
- Acessar diretório projeto
```cd gin-rest-api```
- Executar testes de integração:
```go test```