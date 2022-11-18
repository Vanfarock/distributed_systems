#!/bin/bash

echo "Installing and setting up monitoring tools..."
helm install --wait --timeout 15m \
  --namespace monitoring --create-namespace \
  --repo https://prometheus-community.github.io/helm-charts \
  monitoring kube-prometheus-stack

echo "Port forward Prometheus 9000 -> 9090"
echo "Port forward Grafana 9001 -> 80"
echo ""
kubectl port-forward svc/monitoring-kube-prometheus-prometheus -n monitoring 9000:9090 &\
kubectl port-forward svc/monitoring-grafana -n monitoring 9001:80
