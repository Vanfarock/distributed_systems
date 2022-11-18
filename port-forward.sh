#!/bin/bash

echo "Port forward Tweet-Writer 8000 -> 8080" &\
echo "Port forward Prometheus 9000 -> 9090" &\
echo "Port forward Grafana 9001 -> 80" &\
echo "" &\

kubectl port-forward svc/tweet-writer -n tweet-writer 8000:8080 &\
kubectl port-forward svc/monitoring-kube-prometheus-prometheus -n monitoring 9000:9090 &\
kubectl port-forward svc/monitoring-grafana -n monitoring 9001:80 &\

sleep 1
