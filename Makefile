.PHONY: test-user

# Executa os testes apenas para o pacote de usecases/user
test-user:
	go clean -testcache
	go test -coverprofile=./internal/usecases/user/coverage.out ./internal/usecases/user

user-coverage:
	go tool cover -html=./internal/usecases/user/coverage.out