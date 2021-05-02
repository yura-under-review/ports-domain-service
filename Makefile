
.PHONY: deps
deps:
	go mod tidy


.PHONY: build
build:
	go build -o artifacts/svc .


.PHONY: lint
lint:
	golangci-lint run --allow-parallel-runners -v -c .golangci.yml

.PHONY: gen_proto
gen_proto:
#	go install google.golang.org/protobuf/cmd/protoc-gen-go
#	go install github.com/90poe/service-chassis/protobuf/protoc-gen-gofullmethods
	protoc -I api api/ports-domain-service.proto --go-grpc_out=api --gofullmethods_out=api --go_out=api
