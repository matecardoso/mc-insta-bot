# MC Insta Bot

MC Insta Bot é um sistema desenvolvido em Go utilizando o framework Fiber, seguindo os princípios de Clean Architecture e Clean Code. O objetivo deste projeto é criar um bot para automatizar interações com o Instagram.

## Tecnologias Utilizadas

- **Linguagem:** Go
- **Framework:** Fiber
- **Banco de Dados:** MongoDB
- **Containerização:** Docker e Docker Compose
- **Gerenciamento de Configurações:** godotenv

## Pré-requisitos

- [Go](https://golang.org/dl/) 1.19 ou superior
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Instalação

1. Clone o repositório:

   ```sh
   git clone https://github.com/matecardoso/mc-insta-bot
   cd mc-insta-bot
   ```

2. Crie um arquivo `.env` com suas credenciais do Instagram e a URI do MongoDB:

3. Construa e execute o projeto usando Docker Compose:

   ```sh
   docker-compose up --build
   ```

4. O sistema estará disponível em `http://localhost:3000`.

## Licença

Este projeto está licenciado sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
