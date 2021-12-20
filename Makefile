APP_NAME := stp-exporter
VERSION := v1.0.0

build:
	@go build
	@docker build -t $(APP_NAME):$(VERSION) .

run:
	@docker-compose up -d

stop:
	@docker stop $(APP_NAME)

log:
	@docker logs $(APP_NAME)

terminal:
	@docker exec -it $(APP_NAME) sh
