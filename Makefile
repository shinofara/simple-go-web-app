install-deps:
	@docker run --rm -v $(PWD):/work shinofara/docker-glide:0.12.3 install

## Local
test-all: test vet lint

test:
	@go test $$(glide novendor)

vet:
	@go vet $$(glide novendor)

lint:
	@for pkg in $$(go list ./... | grep -v /vendor/) ; do \
		golint $$pkg ; \
	done

## CI
ci-test:
	@cd "$(WORK_DIR)/src/$(IMPORT_PATH)/" && \
	go test -race -v $$(glide novendor) | go-junit-report -set-exit-code=true > $(CIRCLE_TEST_REPORTS)/golang/junit.xml

ci-vet:
	@cd "$(WORK_DIR)/src/$(IMPORT_PATH)/" && go vet $$(glide novendor)

ci-lint:
	@cd "$(WORK_DIR)/src/$(IMPORT_PATH)/" && \
	for pkg in $$(go list ./... | grep -v /vendor/) ; do \
		golint $$pkg ; \
	done
