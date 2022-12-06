
start:
	docker network inspect currencynetwork >/dev/null 2>&1 || \
            docker network create currencynetwork

	docker compose up -d

stop:
	docker compose down

push:
	docker compose build