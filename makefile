make integration-tests:
	go test ./tests/integration -v

make unit-tests:
	go test ./tests/unit -v

make run:
	go run src/endpoints/endpoints.go