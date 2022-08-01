build-docker:
	docker-compose build
run-docker:
	docker-compose up -d
down-docker:
	docker-compose down
clean-start: build-docker run-docker