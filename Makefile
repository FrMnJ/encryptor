build: 
	@go build -o ./bin/encryptor.exe ./cmd/main.go
run: build
	@./bin/encryptor