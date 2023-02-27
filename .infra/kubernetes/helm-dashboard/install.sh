#!/bin/bash

helm repo add komodorio https://helm-charts.komodor.io
helm upgrade --install helm-dashboard komodorio/helm-dashboard -n helm-dashboard --create-namespace -f values.yaml --debug --atomic