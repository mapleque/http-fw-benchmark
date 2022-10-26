.PHONY: all
all: build run

.PHONY: build
build:
	@docker pull rust:1.64-slim
	@docker build -t rust-actix .

.PHONY: run
run:
	@docker run --rm -d --name rust-actix -p 8080:8080 rust-actix

.PHONY: log
log:
	@docker logs -f rust-actix

.PHONY: clean
clean:
	@docker stop rust-actix
