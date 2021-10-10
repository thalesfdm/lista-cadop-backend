# lista-cadop-backend
Aplicação que, ao ser iniciada, carrega um arquivo **.csv** em memória, fornecendo uma interface de acesso via HTTP, através da qual podem ser realizadas buscas textuais. A API foi desenvolvida para servir a seguinte aplicação: [lista-cadop-frontend](https://github.com/thalesfdm/lista-cadop-frontend).

## Requisitos
- [Git](https://git-scm.com)
- [Go](https://golang.org)

## Instruções
Clonar este repositório e executar os seguintes comandos no diretório raiz do projeto:
```
go mod vendor
go mod download
go run main.go
```
