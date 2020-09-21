# enum-golang
poc using struct and enum in golang

#### Run

```bash
go run cmd/main.go
```

#### Tests

```bash
go test -v ./...
```

#### Coverage

```bash
go test -v ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
```
