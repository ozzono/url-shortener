with-docker:
	docker-compose up -d

without-docker:
	docker-compose up -d mongodb;
	./cmd/url-shortener

tidy:
	go mod tidy;
	go mod download;

mongo-logs:
	docker logs -f mongodb

api-logs:
	docker logs -f golang

rebuild:
	go build -o cmd/url-shortened ./cmd/main.go

rebuild-tests:
	go test ./internal/handler -c
	go test ./internal/database -c

tests:
	./handler.test -test.v
	./database.test -test.v
