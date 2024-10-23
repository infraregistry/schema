include .env

surrealdb/validate:
	@surreal validate surrealdb/*.surql

surrealdb/import: validate
	@for f in $(shell find surrealdb -name '*.surql'); do \
		echo "Importing $$f..."; \
		surreal import --conn http://localhost:9000 --user root --pass root --namespace $(NAMESPACE) --database $(DATABASE) $$f; \
	done

sql/build:
	go run github.com/steebchen/prisma-client-go generate

sql/push: sql/build
	go run github.com/steebchen/prisma-client-go db push