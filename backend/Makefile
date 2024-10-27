.PHONY: build up down

build:
	docker-compose build

up:
	docker-compose up

down:
	docker-compose down

# Для очистки кэша и пересборки
rebuild:
	docker-compose down --rmi all
	docker-compose build --no-cache
	docker-compose up
