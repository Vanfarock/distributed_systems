#!/bin/bash

echo "Setting up shortest-path..."
kubectl create namespace shortest-path
kubectl apply -f ./k8s/shortest_path/deployment.yaml -n shortest-path
kubectl apply -f ./k8s/shortest_path/service.yaml -n shortest-path
kubectl apply -f ./k8s/shortest_path/ingress.yaml -n shortest-path
kubectl apply -f ./k8s/shortest_path/hpa.yaml -n shortest-path

kubectl apply -f ./k8s/metrics-server.yaml

sleep 1