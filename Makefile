cli:
	go install github.com/zoli/emru/cmd/emru
daemond:
	go build -o emrud github.com/zoli/emru/daemon
test:
	go test ./...
