.PHONY: all
all: build run

.PHONY: build
build:
	@docker pull golang:1.18.7
	@docker build -t go-http .

.PHONY: run
run:
	@docker run --rm -d --name go-http -p 8080:8080 go-http

.PHONY: log
log:
	@docker logs -f go-http

.PHONY: clean
clean:
	@docker stop go-http
