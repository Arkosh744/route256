LOCAL_BIN:=$(CURDIR)/bin


install-go-deps:
	mkdir -p bin
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v0.10.1
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2
	GOBIN=$(LOCAL_BIN) go install github.com/rakyll/statik@latest

generate:
	mkdir -p pkg/swagger
	make generate-checkout-api
	make generate-loms-api
	make generate-notifications-api
	make generate-product-api
	mkdir -p pkg/statik
	./bin/statik -src=pkg/swagger -include='*.css,*.html,*.js,*.json,*.png'
	cp -r ./statik/statik.go pkg/statik/statik.go
	rm -rf ./statik

generate-checkout-api:
	mkdir -p pkg/checkout_v1
	mkdir -p pkg/swagger
	protoc --proto_path api/checkout_v1 --proto_path vendor.protogen \
	--go_out=pkg/checkout_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/checkout_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--validate_out lang=go:pkg/checkout_v1 --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=bin/protoc-gen-validate \
	--grpc-gateway_out=pkg/checkout_v1 --grpc-gateway_opt=paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
	--openapiv2_out=allow_merge=true,merge_file_name=api:pkg/swagger \
	--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
	api/checkout_v1/checkout.proto

generate-loms-api:
	mkdir -p pkg/loms_v1
	mkdir -p pkg/swagger
	protoc --proto_path api/loms_v1 --proto_path vendor.protogen \
	--go_out=pkg/loms_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/loms_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--validate_out lang=go:pkg/loms_v1 --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=bin/protoc-gen-validate \
	api/loms_v1/loms.proto

generate-notifications-api:
	mkdir -p pkg/notifications_v1
	protoc --proto_path api/notifications_v1 --proto_path vendor.protogen \
	--go_out=pkg/notifications_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/notifications_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--validate_out lang=go:pkg/notifications_v1 --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=bin/protoc-gen-validate \
	api/notifications_v1/notifications.proto

generate-product-api:
	mkdir -p pkg/product_v1
	mkdir -p pkg/swagger
	protoc --proto_path api/product_v1 --proto_path vendor.protogen \
	--go_out=pkg/product_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/product_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/product_v1/product.proto

lint:
	gofumpt -w -extra .
	golangci-lint run checkout/... loms/... notifications/... libs/... pkg/...

build-all:
	cd checkout && GOOS=linux GOARCH=amd64 make build
	cd loms && GOOS=linux GOARCH=amd64 make build
	cd notifications && GOOS=linux GOARCH=amd64 make build

run-all: build-all
	sudo docker compose up --force-recreate --build -d
	cd checkout && make local-migration-up
	cd loms && make local-migration-up

install-go-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

test:
	go clean -testcache;
	cd checkout && make test-all
	cd loms && make test
	cd notifications && go test -cover ./...

regenerate_mocks:
	cd checkout && go generate -run="mockgen .*" -x ./...
	cd loms && go generate -run="mockgen .*" -x ./...
	cd libs && go generate -run="mockgen .*" -x ./...

.PHONY: logs
logs:
	mkdir -p logs/data
	touch logs/data/checkout.txt
	touch logs/data/loms.txt
	touch logs/data/notifications.txt
	touch logs/data/offsets.yaml
	sudo chmod -R 777 logs/data
	cd logs && sudo docker compose up

precommit:
	cd checkout && make precommit
	cd loms && make precommit
	cd notifications && make precommit

vendor-proto:
		@if [ ! -d vendor.protogen/validate ]; then \
			mkdir -p vendor.protogen/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
			mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
			rm -rf vendor.protogen/protoc-gen-validate ;\
		fi
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
			mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
			git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
			mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
			rm -rf vendor.protogen/openapiv2 ;\
		fi
