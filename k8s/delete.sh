#!/bin/bash

echo "Deleting shortest-path..."
kubectl delete deploy shortest-path -n shortest-path
kubectl delete svc shortest-path -n shortest-path
kubectl delete ingress shortest-path-ingress -n shortest-path

sleep 1