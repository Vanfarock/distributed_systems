#!/bin/bash

echo "Deleting tweet-writer..."
kubectl delete deploy tweet-writer -n tweet-writer
kubectl delete svc tweet-writer -n tweet-writer
kubectl delete ingress tweet-writer-ingress -n tweet-writer

sleep 1