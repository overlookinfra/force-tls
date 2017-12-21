CONTAINER_NAME = puppetlabs/force-tls
CONTAINER_TAG = latest

.DEFAULT_GOAL := compile

.PHONY : compile
compile : build/force-tls

build/force-tls : force-tls.go
	GOOS=linux go build $(GO_ARGS) -o build/force-tls ./force-tls.go

.PHONY : container
container : build/force-tls
	docker build -t $(CONTAINER_NAME):$(CONTAINER_TAG) .

