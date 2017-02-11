setup: clean build-ssl
	go run /usr/local/go/src/crypto/tls/generate_cert.go --host localhost
	glide install

build-ssl:
	$(eval TMPFILE := $(shell mktemp))
	curl https://gist.githubusercontent.com/shinofara/3a5295bd55b2fde4c092f0d16e492473/raw/dc3ba9c93d7547b61560f4683b5240a3cd736507/generate.sh > $(TMPFILE)
	sh $(TMPFILE) -h localhost

clean:
	rm -rf cert.pem key.pem vendor localhost.*

run:
	docker-compose up -d mysql
	go run main.go -ssl-cert=localhost.crt -ssl-key=localhost.pem -http-port=8080
