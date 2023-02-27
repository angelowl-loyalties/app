#!/bin/bash

# helm repo add bitnami https://charts.bitnami.com/bitnami

ETCD_ROOT_PASSWORD=$(\
    sudo kubectl get secret \
    --namespace "angelowl-etcd" \
    etcd \
    -o jsonpath="{.data.etcd-root-password}" \
    | base64 -d \
)

helm upgrade --install etcd bitnami/etcd -f values.yaml -n angelowl-etcd --create-namespace --atomic --debug --set auth.rbac.rootPassword=$ETCD_ROOT_PASSWORD