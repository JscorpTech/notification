up-auth:
	docker compose --profile auth up -d

down-auth:
	docker compose --profile auth down
up-payment:
	docker compose --profile payment up -d
down-payment:
	docker compose --profile payment down

down-all:
	docker compose --profile payment --profile auth down