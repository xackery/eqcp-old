# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
VERSION := v0.0.1
NAME := eqcp

.PHONY: init
init:
	@echo "Priming a fresh database"
	-docker-compose down
	-rm -rf db/
	-mkdir tmp
	wget http://edit.projecteq.net/weekly/peq_beta.zip -O tmp/peq_beta.zip
	cd tmp && unzip -o peq_beta.zip	
	cd tmp && rm drop*.sql *.zip data_tables.sql load_* source_views.sql
	docker-compose up -d
	@docker-compose logs mariadb
	@echo "Wait a bit for database to be injected." 
	@echo "you can run 'docker-compose logs mariadb' to check status of import"
	@echo "host: 127.0.0.1:3306 db: eqemu, user: eqemu, pass: eqemupass"
.PHONY: up
up:
	@# I delete tmp/ so that the initializer scripts don't happen again
	@-rm -rf tmp/
	@docker-compose up -d
.PHONY: down
down:
	@-rm -rf tmp/
	@docker-compose down
.PHONY: logs-mariadb
logs-mariadb:
	@docker-compose logs mariadb
.PHONY: build-all
build-all: proto sanitize
	@echo "Preparing ${NAME} ${VERSION}"
	@rm -rf bin/*
	@-mkdir -p bin/
	@echo "Building Linux"
	@GOOS=linux GOARCH=amd64 go build -ldflags="-X main.Version=${VERSION} -s -w" -o bin/${NAME}-${VERSION}-linux-x64 main.go	
	@GOOS=linux GOARCH=386 go build -ldflags="-X main.Version=${VERSION} -s -w" -o bin/${NAME}-${VERSION}-linux-x86 main.go
	@echo "Building Windows"
	@GOOS=windows GOARCH=amd64 go build -ldflags="-X main.Version=${VERSION} -s -w" -o bin/${NAME}-${VERSION}-win-x64.exe main.go
	@GOOS=windows GOARCH=386 go build -ldflags="-X main.Version=${VERSION} -s -w" -o bin/${NAME}-${VERSION}-win-x86.exe main.go
	@echo "Building OSX"
	@GOOS=darwin GOARCH=amd64 go build -ldflags="-X main.Version=${VERSION} -s -w" -o bin/${NAME}-${VERSION}-osx-x64 main.go
.PHONY: sanitize
sanitize:
	@goimports -w .
	@golint

PROTO_VERSION=3.8.0
GO_PLUGIN=bin/protoc-gen-go
GO_PROTOBUF_REPO=github.com/golang/protobuf
GO_PTYPES_ANY_PKG=$(GO_PROTOBUF_REPO)/ptypes/any
SWAGGER_PLUGIN=bin/protoc-gen-swagger
PROTO_FILES=$(shell find proto -name '*.proto')
PROTO_OUT=/src/pb/
$(GO_PLUGIN):
	dep ensure -vendor-only
	go install ./vendor/$(GO_PLUGIN_PKG)
	go build -o $@ $(GO_PLUGIN_PKG) -ldflags="-s -w -X 'main.Version=${VERSION}'"
proto-clean:
	@echo "removing generated contents..."
	@rm -rf pb/*.pb.*go
	-@rm -rf swagger/*
	@mkdir -p swagger
.PHONY: proto
proto: proto-clean ## Generate protobuf files
	@echo "proto > pb"
	@(docker run --rm -v ${PWD}:/src xackery/protobuf:$(PROTO_VERSION) protoc \
	-I/protobuf/src \
	-I/src \
	-I/grpc \
	-I/grpc/third_party/googleapis \
	$(PROTO_FILES) \
	--grpc-gateway_out=logtostderr=true:$(PROTO_OUT) \
	--swagger_out=logtostderr=true,allow_merge=true:swagger/ \
	--go_out=plugins=grpc+retag:$(PROTO_OUT))
	@(mv pb/proto/* pb/)
	@(rm -rf pb/proto)
	@$(MAKE) sanitize