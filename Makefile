up-user:
	docker compose --profile user up -d
down-user:
	docker compose --profile user down
up-payment:
	docker compose --profile payment up -d
down-payment:
	docker compose --profile payment down
down-all:
	docker compose --profile payment --profile user --profile notification down