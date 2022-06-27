.PHONY: init
init:
	docker build -t goenv .

.PHONY: clean
clean:
	docker volume prune
	docker image prune
	docker container prune

.PHONY: test
test:
	docker run goenv test

.PHONY: build
build:
	docker run -v $(CURDIR)/bin:/go/src/app/bin goenv build -o bin/main