DC=docker-compose
SLACK_NOTIFY=$(DC) -f resources/docker-compose.slack.yml run --rm
build:
	@echo "building nothing"

dev:
	@echo "deploy dev something"

stg:
	@echo "deploy staging something"

prd:
	@echo "deploy prd something"

slack_success:
	$(SLACK_NOTIFY) success

slack_failure:
	$(SLACK_NOTIFY) failure

