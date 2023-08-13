run:
	go run main.go

mock_model:
	go get github.com/golang/mock/mockgen/model

mock:
	mockgen -destination cmd/database/mock/store.go  github.com/durotimicodes/xanda_task_R3_D3/cmd/database/repository SpaceshipRepository

test:
	go test -v -cover ./...

mod:
	go mod tidy
