main:
	docker compose -f docker-compose.yaml -f docker-compose.prod.yaml up -d

web:
	docker compose -f docker-compose.yaml -f docker-compose.prod.yaml up -d --no-deps --build web

pocketbase:
	docker compose -f docker-compose.yaml -f docker-compose.prod.yaml up -d --no-deps --build pocketbase
