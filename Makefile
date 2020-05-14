TAG?=0.3

DOCKER_PROJECT?=runyonsolutions
DOCKER_IMAGE_NAME?=request-log

docker:
	docker build -t ${DOCKER_PROJECT}/${DOCKER_IMAGE_NAME}:${TAG} .

push:
	docker push ${DOCKER_PROJECT}/${DOCKER_IMAGE_NAME}:${TAG}
