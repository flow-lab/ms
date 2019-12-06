SHELL := /bin/bash

SRV_NAME := ms
PROJECT := diatom-ai
HOSTNAME := eu.gcr.io
DOCKER_IMG := flowlab/${SRV_NAME}

deps:
	go mod download

deps-reset:
	git checkout -- go.mod

tidy:
	go mod tidy

verify:
	go mod verify

test:
	go test -mod=readonly -covermode=atomic -v ./...

build-docker:
	docker build -t ${DOCKER_IMG} .

docker-tag:
	docker tag ${DOCKER_IMG} ${HOSTNAME}/${PROJECT}/${SRV_NAME}

docker-push:
	gcloud docker -- push ${HOSTNAME}/${PROJECT}/${SRV_NAME}

docker-clean:
	docker system prune -f

# minikube
minikube-init:
	eval $(minikube docker-env)

minikube-local:
	kubectl apply -f k8-local.yml

minikube-get-pod:
	kubectl get pod

minikube-app-url:
	minikube service ${SRV_NAME} --url
