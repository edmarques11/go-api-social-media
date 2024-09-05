up:
	docker compose up
down:
	docker compose down
test:
	go test ./... -v
test-cover:
	go test ./... -v --cover coverage.txt
go-show-cover:
	go tool cover --html=coverage.txt