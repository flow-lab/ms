SHELL := /bin/bash

SRV_NAME := ms
PROJECT := test
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
	docker tag ${DOCKER_IMG} ${HOSTNAME}/${PROJECT}/${SRV_NxAME}

docker-push:
	gcloud docker -- push ${HOSTNAME}/${PROJECT}/${SRV_NAME}

# minikube -> https://kubernetes.io/docs/setup/learning-environment/minikube
#minikube-create:
#	kubectl create deployment ${SRV_NAME} --image=${HOSTNAME}/${PROJECT}/${SRV_NAME}:latest
#
#minikube-delete:
#	kubectl delete deployment ${SRV_NAME}
#
#minikube-expose:
#	kubectl expose deployment ${SRV_NAME} --type=NodePort --port=8080
#
#minikube-get-pod:
#	kubectl get pod
#
#minikube-service:
#	minikube service ${SRV_NAME} --url