.PHONY: dev
dev:
	docker-compose -f development/docker-compose.yml up --build

.PHONE: prod
prod:
	docker-compose -f docker-compose.yml up --build