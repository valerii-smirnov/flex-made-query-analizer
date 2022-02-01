current_dir = $(shell pwd)

run:
	docker-compose -f docker/docker-compose.yaml up -d

stop:
	docker-compose -f docker/docker-compose.yaml down -v

test:
	docker run -v $(current_dir):/app -w /app golang:1.17 go test ./...  \
		-coverpkg=./application/... \
		-coverprofile ./coverage.out && go tool cover -func ./coverage.out

test-local:
	go test ./...  \
    		-coverpkg=./application/... \
    		-coverprofile ./coverage.out && go tool cover -func ./coverage.out
