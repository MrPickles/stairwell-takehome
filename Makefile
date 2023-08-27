.PHONY: build  deploy
IMAGE_NAME := stairwell-takehome-entrypoint:latest
DOCKER_HUB_USERNAME := guru

build:
	@echo "Building container and pushing it in local docker image repository"
	@docker build -t $(IMAGE_NAME) entrypoint/.

deploy: build
	@echo "Pushing image to dockerhub"
	@docker tag $(IMAGE_NAME) $(DOCKER_HUB_USERNAME)/$(IMAGE_NAME)
	@docker push $(DOCKER_HUB_USERNAME)/$(IMAGE_NAME)
