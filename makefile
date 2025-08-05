make integration-tests:
	go test ./tests/integration -v
	
make docker: 
	docker build -t mein-app .
	docker run -p 8080:8080 mein-app
	
make docker-rm:
	docker stop mein-app
	docker rm mein-app

make unit-tests:
	go test ./tests/unit -v

make run:
	go run run.go