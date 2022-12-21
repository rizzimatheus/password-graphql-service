# Senha Válida - Studio Sol

## How to Run

`URL: http://localhost:8080/graphql`

### With Docker

```bash
docker build -t verifier-graphql-service .
```

```bash
docker run --name verifier-graphql -p 8080:8080 -d verifier-graphql-service
```

### With Go

```bash
go run server.go
```

#### Tests

```bash
go test -v ./cmd/api
```
