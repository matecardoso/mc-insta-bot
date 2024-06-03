# MC Insta Bot

MC Insta Bot é um sistema desenvolvido em Go utilizando o framework Fiber, seguindo os princípios de Clean Architecture e Clean Code. O objetivo deste projeto é criar um bot para automatizar interações com o Instagram.

## Tecnologias Utilizadas

- **Linguagem:** Go
- **Framework:** Fiber
- **Banco de Dados:** MongoDB
- **Containerização:** Docker e Docker Compose
- **Gerenciamento de Configurações:** godotenv
- **Hot-reload:** air

## Pré-requisitos

- [Go](https://golang.org/dl/) 1.19 ou superior
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [air](https://github.com/cosmtrek/air) (já incluído no Dockerfile de desenvolvimento)

## Instalação

1. Clone o repositório:

   ```sh
   git clone https://github.com/matecardoso/mc-insta-bot
   cd mc-insta-bot
   ```

2. Crie um arquivo `.env` com suas credenciais do Instagram e a URI do MongoDB:

   ```
   INSTAGRAM_USERNAME=seu_usuario
   INSTAGRAM_PASSWORD=sua_senha
   MONGO_URI=sua_mongo_uri
   ```

3. Construa e execute o projeto usando Docker Compose:

   ```sh
   make dev
   ```

4. O sistema estará disponível em `http://localhost:3000`.

## Comandos do Makefile

- **Construir o serviço Docker:**
  ```sh
  make build
  ```

- **Subir o serviço Docker em modo destacado:**
  ```sh
  make up
  ```

- **Derrubar o serviço Docker:**
  ```sh
  make down
  ```

- **Mostrar logs do serviço Docker:**
  ```sh
  make logs
  ```

- **Desenvolvimento com hot-reload:**
  ```sh
  make dev
  ```

- **Executar os testes Go dentro do contêiner Docker:**
  ```sh
  make test
  ```

- **Limpar todos os contêineres, volumes, imagens e órfãos do Docker:**
  ```sh
  make clean
  ```

## Licença

Este projeto está licenciado sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
