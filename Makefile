build:
	docker compose build
run:
	docker compose up
init:
	docker build -t configurator-app -f /configurator/configurator.dockerfile .
	docker run configurator-app

