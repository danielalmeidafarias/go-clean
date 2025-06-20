.PHONY: test-user

# Executa os testes apenas para o pacote de usecases/user
test-user:
	go clean -testcache
	go test -v -coverprofile=./internal/usecases/user/coverage.out ./internal/usecases/user

user-coverage:
	go tool cover -html=./internal/usecases/user/coverage.out

test-task:
	go clean -testcache
	go test -v -coverprofile=./internal/usecases/task/coverage.out ./internal/usecases/task

task-coverage:
	go tool cover -html=./internal/usecases/task/coverage.out

test: test-user test-task

coverage: user-coverage task-coverage