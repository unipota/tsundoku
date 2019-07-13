.PHONY: dev
dev:
	docker-compose -f ../development/docker-compose.yml up --build

.PHONY: setup
setup:
	go mod download

.PHONY: build
build: 
	go build . -o app