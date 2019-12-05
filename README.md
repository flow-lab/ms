# Microsevice

Build in docker golang microservice. Check Makefile for all important stuff.

## Local testing with Minikube

For local Kubernetes testing it can be helpfull to use [Minikube](https://kubernetes.io/docs/setup/learning-environment/minikube/).

To deploy locally to minikube:

### Start minikube
`minikube start`

### Set docker env
`eval $(minikube docker-env)`

### Build docker image, deploy to minikube, and get url to service
Given that your minikube is already running. Run:

`make build-docker minikube-local minikube-app-url`
 