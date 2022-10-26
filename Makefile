all: install run

.PHONY: install
install: bin/tool
	@cd go-http && make build
	@cd rust-actix && make build
	@cd nginx && make build

bin/tool: tool/main.go
	@-mkdir -p bin
	@cd tool && go build -o ../bin/tool main.go

.PHONY: run
run:
	@cd go-http && make run
	@bin/tool
	@cd go-http && make clean
	@cd rust-actix && make run
	@bin/tool
	@cd rust-actix && make clean
	@cd nginx && make run
	@bin/tool
	@cd nginx && make clean

.PHONY: clean
clean:
	@-rm -rf bin
