GO ?= GO111MODULE=on go

export CGO_ENABLED=0

# 设置 操作系统的类型
export GOOS=linux

# 设置 cpu 架构
export GOARCH=amd64

.PHONY: bin

bin:
	@make tidy
	@$(GO) build -o bin/app cmd/main.go

run:
	@make tidy
	@$(GO) run cmd/main.go

tidy:
	@$(GO) mod tidy

image:
	@make bin
	@docker build --force-rm --compress -q -t onlyslient/gifl:0.1 .