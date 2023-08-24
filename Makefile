BINARY_NAME=subscription-service
POSTGRES_DSN="user=postgres password=postgres dbname=subscriptions host=localhost port=5432 sslmode=disable search_path=subscriptions"
REDIS_URI="127.0.0.1:6379"
SERVER_PORT=3000

build:
	@echo "Building..."
	env CGO_ENABLED=0 go build -ldflags="-s -w" -o ${BINARY_NAME} ./
	@echo "Build has finished"

run: build
	@echo "Starting application..."
	@env SERVER_PORT=${SERVER_PORT} POSTGRES_DSN=${POSTGRES_DSN} REDIS_URI=${REDIS_URI} ./${BINARY_NAME} &
	@echo "Application started"

clean:
	@echo "Cleaning application..."
	@go clean
	@echo "Cleaned"

start: run

stop:
	@echo "Stopping application..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Application stopped"

restart: stop start

migrations-up:
	@echo "Running migrations up..."
	@migrate -source file://drivers/postgres/migrations -database postgres://postgres:postgres@localhost:5432/subscriptions?sslmode=disable up
	@echo "Migrations run successfully"

migrations-down:
	@echo "Running migrations down..."
	@migrate -source file://drivers/postgres/migrations -database postgres://postgres:postgres@localhost:5432/subscriptions?sslmode=disable down
	@echo "Migrations run successfully"

test:
	go test -v ./...
