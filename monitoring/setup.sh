#!/bin/bash

echo "Installing and setting up monitoring tools..."
helm install --wait --timeout 15m \
  --namespace monitoring --create-namespace \
  --repo https://prometheus-community.github.io/helm-charts \
  monitoring kube-prometheus-stack
