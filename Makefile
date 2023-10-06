run:
	go run *.go
prep:
	go install go.uber.org/mock/mockgen@latest
init: 
	go mod tidy
mock:
	mockgen -source=operation.go -destination=operation_mock.go -package=main
test:
	go test -race -short -coverprofile=./cov.out operation_test.go operation.go ops_advance.go ops_basic.go