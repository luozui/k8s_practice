#!/bin/bash

kubectl $1 -f app1/ingress.yaml
kubectl $1 -f app1/frontend/app1-frontend-server.yaml
kubectl $1 -f app1/frontend/app1-frontend.yaml
kubectl $1 -f app1/serverend/app1-serverend-server.yaml
kubectl $1 -f app1/serverend/app1-serverend.yaml