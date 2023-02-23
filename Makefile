.DEFAULT_GOAL := run

run: run-server
build: build-server

build-server:
	docker-compose build server

run-server: build-server
	docker-compose up -d server
	docker-compose logs -f server

down:
	docker-compose down

gcloud-init:
	gcloud config set project dev-playground-378601

gcloud-docker-init:
	gcloud auth configure-docker us-east1-docker.pkg.dev

image-push: gcloud-init gcloud-docker-init
	docker buildx build --platform linux/amd64 -t us-east1-docker.pkg.dev/dev-playground-378601/r1/gqlgo-server:dev --push server
	# docker-compose build server && docker tag gqlgo-server us-east1-docker.pkg.dev/dev-playground-378601/r1/gqlgo-server:dev && \
	# docker push us-east1-docker.pkg.dev/dev-playground-378601/r1/gqlgo-server:dev

deploy: image-push
	gcloud run deploy gqlgo-server --image us-east1-docker.pkg.dev/dev-playground-378601/r1/gqlgo-server:dev --platform managed --region us-east1 --allow-unauthenticated --port 8000

.PHONY: build build-server run run-server down gcloud-init gcloud-docker-init image-push
