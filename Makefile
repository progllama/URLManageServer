
.PHONY: test
test:
	@docker-compose -f docker/dev/docker-compose.yml up -d

.PHONY: build
build:
	go build src/cmd/runapp/main.go user_api.go user_api_test.go session_controller_test.go sessions_controller.go urls_controller.go urls_controller_test.go user_controller_test.go users_controller.go login_form.go url_form.go user_form.go favicon.go login_require_middleware.go login_require_middleware_test.go link.go url_dict.go user.go url_dict_repository.go url_repository.go user_repository.go user_repository_mock.go mem_session.go redis_sessoin.go session.go user.go main.go config.go db_config.go redis_config.go server_config.go db.go dsn.go server.go