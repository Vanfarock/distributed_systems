#!/bin/bash

echo "Setting cluster up..."
kind create cluster --name twitter --config kind-config.yaml
