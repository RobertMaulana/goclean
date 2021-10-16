BINARY=engine
test:
	go test -v -cover -covermode=atomic ./...

engine:
	go build -o ${BINARY} cmd/web/*.go

run:
	docker-compose up --build -d

stop:
	docker-compose down