IMG_NAME = bootstrap
ACCOUNT_JJ=164382793440
DOCKER_PUSH_REPOSITORY=dkr.ecr.us-west-2.amazonaws.com
TAG?=latest
IMAGE=$(ACCOUNT_JJ).$(DOCKER_PUSH_REPOSITORY)/$(IMG_NAME)


ci-pr: build-image push-image
ci-release: build-image push-image

aws-iam:
	wget -O tmp/aws-iam-authenticator https://amazon-eks.s3.us-west-2.amazonaws.com/1.14.9/2020-01-22/bin/linux/amd64/aws-iam-authenticator
	chmod +x tmp/aws-iam-authenticator

build-image:
	docker build -t $(IMG_NAME) .
push-image:
	docker tag $(IMG_NAME) $(IMAGE):$(TAG)
	docker push $(IMAGE):$(TAG)