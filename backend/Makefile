# Makefile

# OS에 make 설치 필요
# [install make at ubuntu]
# sudo apt update
# sudo apt install make

# make generate
generate: sqlc mockery

sqlc : 
	sqlc generate
mockery :
	mockery

# make test
test:
	go test ./internal/services/ ./internal/handlers/ -v

# make run
run:
	go run ./cmd/main.go