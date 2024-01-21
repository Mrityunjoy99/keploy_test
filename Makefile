docker-build:
	docker build -t sample_project:local .

docker-svc-up:
	docker-compose up -d

run:
	go run cmd/main.go start

migrate:
	go run cmd/main.go --migrate-all migrate

rollback-last:
	go run cmd/main.go --rollback-last migrate

generate-migration:
	sh ./infra/migration/migration_generator.sh ${name}

test-run:
	GIN_MODE=release go test -race $(VERB) $(MOD)

# Run test with coverage
test-cover:
	GIN_MODE=release CGO_ENABLED=1 go test -race -coverprofile=$(COVER) $(VERB) $(MOD)

test-run-cover: test-cover
	scoverage -package=github.com/sample_project -minCov=85

html-coverage-report: test-cover
	go tool cover -html=$(COVER)
