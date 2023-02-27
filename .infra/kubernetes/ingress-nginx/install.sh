#!/bin/bash

helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm upgrade --install nginx-ingress ingress-nginx/ingress-nginx -n nginx-ingress -f values.yaml --debug --atomic