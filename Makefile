run:
	go run main.go operation.go ops_advance.go ops_basic.go
setup:
	go install go.uber.org/mock/mockgen@latest
init: 
	go mod tidy
mock:
	mockgen -source=operation.go -destination=operation_mock.go -package=main
test:
	go test -race -short -coverprofile=./cov.out ./...

# make and exec binary
build:
	CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix cgo -o go-calculator .
exec: 
	./go-calculator