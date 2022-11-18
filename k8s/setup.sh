#!/bin/bash

echo "Setting up tweet-writer..."
# kubectl create namespace tweet-writer
kubectl apply -f ./k8s/tweet-writer/deployment.yaml # -n tweet-writer
kubectl apply -f ./k8s/tweet-writer/service.yaml # -n tweet-writer
kubectl apply -f ./k8s/tweet-writer/ingress.yaml # -n tweet-writer

sleep 1