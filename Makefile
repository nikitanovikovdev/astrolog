include .env

## Rebuild image and start all needed containers
docker-up:
	docker-compose up -d --build

docker-down:
	docker-compose down