validate:
	@surreal validate surrealdb/*.surql

import: validate
	@for f in $(shell find surrealdb -name '*.surql'); do \
		echo "Importing $$f..."; \
		surreal import --conn http://localhost:9000 --user root --pass root --namespace $(NAMESPACE) --database $(DATABASE) $$f; \
	done