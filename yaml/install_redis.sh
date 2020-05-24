#!/bin/bash

kubectl $1 -f redis/ConfigMap.yaml
kubectl $1 -f redis/Service.yaml
kubectl $1 -f redis/StatefulSet.yaml