postgresinit:
	docker run --name go-chat-db -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

postgres:
	docker exec -it go-chat-db psql

createdb:
	docker exec -it go-chat-db createdb --username=root --owner=root go-chat

dropdb:
	docker exec -it go-chat-db dropdb go-chat

.PHONY: postgresinit postgres createdb dropdb
