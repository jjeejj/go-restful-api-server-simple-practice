all: tool
	@go build -v .
clean:
	rm -f go-restful-api-server-simple-practice
tool:
	gofmt -w .
	go vet . | grep -v verdor;true
ca:
	openssl req -new -nodes -keyout conf/server.key -out conf/server.crt -days 3650 -x509 \
	-subj '/C=CN/ST=Beijing\L=Beiing/O=Company/OU=IT/CN=127.0.0.1/emailAddress=1048911681@qq.com'
help:
	@echo "make - compile the source code"
	@echo "make clean - remobe binary flle and vim swap files"
	@echo "make gotool - run go tool 'fmt' and 'vet'"
	@echo "make ca - generate ca files"

.PHONY: clean gotool ca help