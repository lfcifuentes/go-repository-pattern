### Repository Pattern Example

### Run api
```shell
go run main.go run
```

### Database Schema
```shell
go run main.go migrate
```

#### Testing

```bash
go test -cover ./...
```
```bash
go clean -testcache 
```
```bash
go test  ./... -coverprofile=coverage.out
``` 
```bash
go tool cover -html=coverage.out
```

### Start Mysql Services
```shell
docker-compose up -d
```

### Delete
```shell
docker-compose down
```