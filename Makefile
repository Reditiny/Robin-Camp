docker-up:
	docker-compose -f ./deploy/docker-compose.yml up --build --wait

docker-down:
	docker-compose -f ./deploy/docker-compose.yml down

test-e2e: docker-up
	./e2e-test.sh
	make docker-down