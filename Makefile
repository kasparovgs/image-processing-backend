.PHONY: launch_services launch_with_tests stop_services build_services

launch_services: build_services migrate
	docker compose up

launch_with_tests: build_services migrate
	docker compose --profile test up --abort-on-container-exit --exit-code-from app_test

stop_services:
	docker compose --profile test down

build_services:
	docker compose build --no-cache

migrate:
	docker compose run --rm migrate