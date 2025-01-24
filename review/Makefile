docker-build:
	docker build -t schwarz-go-test .
.PHONY: docker-build

docker-up:
	docker run --name schwarz-go-test -p 8080:8080 --rm -it -v $(PWD):/app schwarz-go-test
.PHONY: docker-up

docker-down:
	docker stop schwarz-go-test
.PHONY: docker-down

run:
	go run cmd/coupon_service/main.go
.PHONY: run

lint:
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.63.4 golangci-lint run
.PHONY: lint

test:
	go test	./... -v -cover
.PHONY: test
