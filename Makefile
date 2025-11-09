build-dev-image:
	docker build -t go-cloud -f Dockerfile .	

run-dev-container:
	docker run -d -p 8080:8080 --name go-cloud-app go-cloud

run-access-container:
	docker exec -it go-cloud-app sh

run-tests:
	go test ./... 

run-compose-up:
	docker-compose up -d
