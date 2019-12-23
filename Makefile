#!/usr/bin/make


REGISTRY_HOST = registry.freedomcore.ru
REGISTRY_NAMESPACE = freedomcore
REGISTRY_IMAGE = micro-database
REGISTRY_PATH = $(REGISTRY_NAMESPACE)/$(REGISTRY_IMAGE)
IMAGES_PREFIX := local

PUBLISH_TAGS = latest
PULL_TAG = latest

MICRO_IMAGE = $(REGISTRY_HOST)/$(REGISTRY_PATH)
MICRO_IMAGE_LOCAL_TAG = $(IMAGES_PREFIX)_$(REGISTRY_IMAGE)
MICRO_IMAGE_DOCKERFILE = Dockerfile
MICRO_IMAGE_CONTEXT = .


docker_bin := docker
docker_compose_bin := $(shell command -v docker-compose 2> /dev/null)

all_images =	$(MICRO_IMAGE) \
				$(MICRO_IMAGE_LOCAL_TAG)

.PHONY: help pull build push login test clean \
		back-pull back back-push back-rp
.DEFAULT_GOAL := help

help: ## Show this help
	echo "Allowed for overriding next properties:\
    	    PULL_TAG - Tag for pulling images before building own\n\
    	              ('latest' by default)\n\
    	    PUBLISH_TAGS - Tags list for building and pushing into remote registry\n\
    	                   (delimiter - single space, 'latest' by default)\n\n\
    	  Usage example:\n\
    	    make PULL_TAG='v1.2.3' PUBLISH_TAGS='latest v1.2.3 test-tag' app-push"

# --- [ Frontend ] -------------------------------------------------------------------------------------------------------

back-pull: ## Frontend - pull latest Docker image (from remote registry)
	-$(docker_bin) pull "$(MICRO_IMAGE):$(PULL_TAG)"

back: back-pull ## Frontend - build Docker image locally
	$(docker_bin) build \
	  --cache-from "$(MICRO_IMAGE):$(PULL_TAG)" \
	  --tag "$(MICRO_IMAGE_LOCAL_TAG)" \
	  -f $(MICRO_IMAGE_DOCKERFILE) $(MICRO_IMAGE_CONTEXT)

back-push: back-pull ## Frontend - tag and push Docker image into remote registry
	$(docker_bin) build \
	  --cache-from "$(MICRO_IMAGE):$(PULL_TAG)" \
	  $(foreach tag_name,$(PUBLISH_TAGS),--tag "$(MICRO_IMAGE):$(tag_name)") \
	  -f $(MICRO_IMAGE_DOCKERFILE) $(MICRO_IMAGE_CONTEXT);
	$(foreach tag_name,$(PUBLISH_TAGS),$(docker_bin) push "$(MICRO_IMAGE):$(tag_name)";)

back-rp:
	$(docker_bin) build \
	  --cache-from "$(MICRO_IMAGE):$(PULL_TAG)" \
	  $(foreach tag_name,$(PUBLISH_TAGS),--tag "$(MICRO_IMAGE):$(tag_name)") \
	  -f $(MICRO_IMAGE_DOCKERFILE) $(MICRO_IMAGE_CONTEXT)
	$(foreach tag_name,$(PUBLISH_TAGS),$(docker_bin) push "$(MICRO_IMAGE):$(tag_name)")

# ---------------------------------------------------------------------------------------------------------------------

pull: back-pull ## Pull all Docker images (from remote registry)

build: back ## Build all Docker images

push: back-push ## Tag and push all Docker images into remote registry

login: ## Log in to a remote Docker registry
	@echo $(docker_login_hint)
	$(docker_bin) login $(REGISTRY_HOST)
