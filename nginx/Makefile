.PHONY: all
all: build run

.PHONY: build
build:
	@docker pull nginx:1.23-alpine
	@docker build -t nginx-benchmark .

.PHONY: run
run:
	@docker run --rm -d --name nginx-benchmark -p 8080:8080 nginx-benchmark

.PHONY: log
log:
	@docker logs -f nginx-benchmark

.PHONY: clean
clean:
	@docker stop nginx-benchmark
