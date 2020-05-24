#!/bin/bash

kubectl $1 -f mysql/mysql-pv.yamll
kubectl $1 -f mysql/mysql-deployment.yaml