
# Install all dependencies required.
#
# NOTE: Docker & Docker Compose should already be installed.
.PHONY: install
install:
		curl https://glide.sh/get | sh
		glide install

# Sets up development environment requirements through Docker Compose.
.PHONY: infra
infra:
		cd  deployments/docker/ && \
		docker-compose up -d
		./deployments/docker/wait-docker.sh

# Runs test suite with all development environment requirements.
.PHONY: test
test: infra
		-go test -v -race -cover -timeout=120s $$(glide novendor)
		docker-compose -f deployments/docker/docker-compose.yml down
