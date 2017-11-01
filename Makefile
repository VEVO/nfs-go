DC=docker-compose -f resources/docker-compose.yml
build:
	@echo "building nothing"

dev:
	@echo "deploy dev something"

stg:
	@echo "deploy staging something"

prd:
	@echo "deploy prd something"

slack_success:
	$(DC) run --rm success

slack_failure:
	$(DC) run --rm failure

