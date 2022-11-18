#!/bin/bash

kind create cluster --name twitter --config ./kind/kind-config.yaml

echo ""
echo "Deploying NGINX..."
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml