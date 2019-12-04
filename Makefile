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
	go test -mod=readonly -v ./...

build-docker:
	docker build -t ${DOCKER_IMG} .

build-app: test
	GOOS=linux GOARCH=amd64 go build -mod=readonly -ldflags="-w -s" -o /go/bin/app ./...

run-local:
	docker run -it -p 8080:8080 --rm ${DOCKER_IMG}

docker-tag:
	docker tag ${DOCKER_IMG} ${HOSTNAME}/${PROJECT}/${SRV_NAME}

docker-push:
	gcloud docker -- push ${HOSTNAME}/${PROJECT}/${SRV_NAME}

# minikube
minikube-run:
	kubectl run ${SRV_NAME} --generator=run-pod/v1 --image=${DOCKER_IMG} --image-pull-policy=Never

minikube-delete:
	kubectl delete pod ${SRV_NAME}

minikube-expose:
	kubectl expose pod ${SRV_NAME} --port=8080 --name=${SRV_NAME} --type=NodePort

minikube-get-pod:
	kubectl get pod

minikube-service-url:
	minikube service ${SRV_NAME} --url
