.PHONY: gen-oapi
gen-oapi:
	set -eux
	pnpm --package=@redocly/cli@1.34 dlx \
		redocly bundle openapi/openapi.yml \
		--output openapi/openapi.gen.yml --ext yml
	go tool oapi-codegen -config oapi-codegen.yml openapi/openapi.gen.yml

.PHONY: run
run:
	go run cmd/main.go
