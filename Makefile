SERVICES ?= -f order/docker-compose.yml -f user/docker-compose.yml -f payment/docker-compose.yml -f notification/docker-compose.yml
NAME ?= lamenu


logs:
	docker compose -f $(app)/docker-compose.yml logs -f $(cn)

sh:
	docker compose -f $(app)/docker-compose.yml exec $(cn) sh

up:
	docker compose -f docker-compose.yml up -d
	docker compose -f $(app)/docker-compose.yml up -d

build:
	docker compose  -f $(app)/docker-compose.yml build

down:
	docker compose -f $(app)/docker-compose.yml down

up-all:
	docker compose -p $(NAME) $(SERVICES) up -d

down-all:
	docker compose -$(NAME) $(SERVICES) down