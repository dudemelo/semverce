all: install

install:
	@echo "Installing the dependencies..."
	@go get github.com/pressly/chi
	@echo "Building the project..."
	@go install semverce

run:
	@echo "Running the Go server at http://127.0.0.1:4200"
	@./bin/semverce

clean:
	go clean