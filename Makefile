setup: clean build-ssl
	#go run /usr/local/go/src/crypto/tls/generate_cert.go --host localhost
	glide install
#	go get -u github.com/golang/dep/...
#	dep ensure

build-ssl:
	$(eval TMPFILE := $(shell mktemp))
	curl https://gist.githubusercontent.com/shinofara/3a5295bd55b2fde4c092f0d16e492473/raw/dc3ba9c93d7547b61560f4683b5240a3cd736507/generate.sh > $(TMPFILE)
	sh $(TMPFILE) -h localhost
	mv localhost.* docker/app/config/

clean:
	rm -rf cert.pem key.pem vendor localhost.*

run:
	docker-compose up -d mysql
	docker-compose up

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
	cd "$(WORK_DIR)/src/$(IMPORT_PATH)/" && \
	go test -race -v $$(glide novendor) | go-junit-report -set-exit-code=true > $(CIRCLE_TEST_REPORTS)/golang/junit.xml

ci-vet:
	cd "$(WORK_DIR)/src/$(IMPORT_PATH)/" && go vet $$(glide novendor)

ci-lint:
	cd "$(WORK_DIR)/src/$(IMPORT_PATH)/" && \
	for pkg in $$(go list ./... | grep -v /vendor/) ; do \
		golint $$pkg ; \
	done
