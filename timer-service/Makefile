all: clean install lint test

clean:
	rm -rf ./bin

install:
	GO111MODULE=on go mod verify
	GO111MODULE=on go mod tidy
	rm -rf vendor
	GO111MODULE=on go mod vendor

mocks:
	@echo "mocks regenerating...\n" 
	@go generate -x ./...

example:
	.project/scripts/example.sh $(ARGS)

lint:
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	$$GOPATH/bin/golangci-lint run

migrate:
	go get -u -d github.com/lib/pq
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	$$GOPATH/bin/migrate -path .project/migrations -database $$DSN up


unit_test:
	go test ./... -tags=unit -count=1 -race

integration_test:
	docker-compose -f deployments/docker-compose-postgresql-integration.yml run integration_test

test: unit_test integration_test

coverage:
	echo "" > coverage.txt
	for d in $(go list ./... | grep -v vendor); do
		go test -race -coverprofile=profile.out -covermode=atomic $d
		if [ -f profile.out ]; then
			cat profile.out >> coverage.txt
			rm profile.out
		fi
	done
