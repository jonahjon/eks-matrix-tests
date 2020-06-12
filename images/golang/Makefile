IMG_NAME = golang
ACCOUNT_JJ=164382793440
DOCKER_PUSH_REPOSITORY=dkr.ecr.us-west-2.amazonaws.com
TAG?=go1.14
TAG_13?=go1.13
IMAGE=$(ACCOUNT_JJ).$(DOCKER_PUSH_REPOSITORY)/$(IMG_NAME)


ci-pr: build-image push-image
ci-release: build-image push-image
ci-release-13: build-image-13 push-image-13

build-image:
	docker build -t $(IMG_NAME) .
push-image:
	export $$(aws ecr get-login-password --region us-west-2)
	docker tag $(IMG_NAME) $(IMAGE):$(TAG)
	docker push $(IMAGE):$(TAG)
build-image-13:
	docker build -f Dockerfile-13 -t $(IMG_NAME) .
push-image-13:
	export $$(aws ecr get-login-password --region us-west-2)
	docker tag $(IMG_NAME) $(IMAGE):$(TAG_13)
	docker push $(IMAGE):$(TAG_13)