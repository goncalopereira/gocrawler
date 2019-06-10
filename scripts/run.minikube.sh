#!/usr/bin/env bash
#minikube start
#minikube addons enable freshpod
eval '$(minikube docker-env)'
docker build -t crawler . -f build/Dockerfile
kubectl replace --force -f build/deployment.yaml
kubectl replace --force -f build/service.yaml
kubectl get pods -o=wide
kubectl get deployments -o=wide
kubectl get services -o=wide
#minikube service crawler --url
#kubectl describe pod
IP="$(minikube ip)"
echo "curl --data-urlencode \"url=https://www.goncalopereira.com\" http://${IP}:30000"
