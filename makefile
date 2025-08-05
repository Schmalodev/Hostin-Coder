make integration-tests:
	go test ./tests/integration

make unit-tests:
	go test ./tests/unit

make run:
	go run src/endpoints/endpoints.go