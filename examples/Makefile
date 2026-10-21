.PHONY: test

# Roda os testes de todos os módulos com detector de race (-race).
test:
	@set -e; \
	for dir in $$(find . -name go.mod -exec dirname {} \; | sort); do \
		echo "==> $$dir"; \
		(cd "$$dir" && go test -race ./...); \
	done
