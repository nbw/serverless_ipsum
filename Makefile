build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/ipsum_web web/main.go

.PHONY: clean
clean:
	rm -rf ./bin ./vendor Gopkg.lock

.PHONY: deploy
deploy: clean build
	sls deploy --verbose

get:
	@echo ">> Getting any missing dependencies.."
	go get -t ./...
.PHONY: get

test:
	go test ./...
.PHONY: test

fmt:
	@echo ">> Running Gofmt.."

