# Senha Válida - Studio Sol

## How to Run

`URL: http://localhost:8080/graphql`

### With Docker

```bash
docker build -t verifier-service .
```

```bash
docker run --name verifier-container -p 8080:8080 -d verifier-service
```

### With Go

```bash
go run server.go
```

#### Tests

```bash
go test -v ./cmd/api
```
