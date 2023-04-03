#!/bin/bash

helm upgrade --install \
  cassandra cassandra \
  -f values.yaml \
  -n angelowl-cassandra \
  --repo https://charts.bitnami.com/bitnami \
  --timeout 25m0s \
  --create-namespace \
  --atomic --debug