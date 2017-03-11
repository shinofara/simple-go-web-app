IMPORT_PATH=github.com/shinofara/simple-go-web-app

default: install-deps test vet

install-deps:
	@docker run --rm -v $(PWD):/go/src/$(IMPORT_PATH) -w /go/src/$(IMPORT_PATH) supinf/go-dep ensure

update-deps:
	@docker run --rm -v $(PWD):/go/src/$(IMPORT_PATH) -w /go/src/$(IMPORT_PATH) supinf/go-dep ensure -update

## Local
test-all: test vet lint

test:
	@go test $$(go list ./... | grep -v /vendor/)

vet:
	@go vet $$(go list ./... | grep -v /vendor/)

lint:
	@for pkg in $$(go list ./... | grep -v /vendor/) ; do \
		golint $$pkg ; \
	done

## CI
ci-test:
	@cd "$(WORK_DIR)/src/$(IMPORT_PATH)/" && \
	go test -race -v $$(go list ./... | grep -v /vendor/) | go-junit-report -set-exit-code=true > $(CIRCLE_TEST_REPORTS)/golang/junit.xml

ci-vet:
	@cd "$(WORK_DIR)/src/$(IMPORT_PATH)/" && go vet $$(go list ./... | grep -v /vendor/)

ci-lint:
	@cd "$(WORK_DIR)/src/$(IMPORT_PATH)/" && \
	for pkg in $$(go list ./... | grep -v /vendor/) ; do \
		golint $$pkg ; \
	done

ci-test-build:
	go build cmd/example_app/main.go
