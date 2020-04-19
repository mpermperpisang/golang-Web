all: kill run

package:
	@go get github.com/rimiti/kill-port
	@echo "Package installed"

kill:
	@kill-port 8081
	@echo "Port 8081 is killed"

run:
	@echo "Server 8081 is running"
	@go run main.go
