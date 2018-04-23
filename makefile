DOCKER_SLUG=soupstore/git-pull

# Cross compilation
build-docker:
	docker build -t $(DOCKER_SLUG):dev .
