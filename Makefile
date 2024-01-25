default: run

.PHONY: run
run:
	make clean
	make migrate
	go run cmd/server/main.go

.PHONY: build
build:
	go build -C cmd/server -o ../../build/

.PHONY: migrate
migrate:
	sqlite3 data.db < sql/schema.sql

.PHONY: clean
clean:
	-rm -r build/
	-rm data.db
