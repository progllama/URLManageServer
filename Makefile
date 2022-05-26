
.PHONY: up
up:
	@docker-compose up -d

.PHONY: down
down:
	@docker-compose down

.PHONY: prunei
prunei:
	@docker image prune

.PHONY: prunec
prunec:
	@docker container prune

.PHONY: rmi
rmi:
	@docker rmi `docker images -q`

.PHONY: rmc
rmc:
	@docker rm -f `docker ps -a -q`

.PHONY: prunev
prunev:
	docker volume prune

.PHONY: test
test:
	@docker-compose -f docker/dev/docker-compose.yml up -d
