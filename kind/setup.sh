#!/bin/bash

echo "Setting cluster up..."
kind create cluster --name twitter --config ./kind/kind-config.yaml