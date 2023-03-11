#!/bin/bash

# This script is used to install the charts into the cluster.

helm list -n angelowl
helm upgrade --install angelowl-campaignex . -n angelowl --atomic --debug -f ./values.yaml,./values.campaignex.yaml
helm upgrade --install angelowl-informer . -n angelowl --atomic --debug -f ./values.yaml,./values.informer.yaml
helm upgrade --install angelowl-profiler . -n angelowl --atomic --debug -f ./values.yaml,./values.profiler.yaml
helm upgrade --install angelowl-rewarder . -n angelowl --atomic --debug -f ./values.yaml,./values.rewarder.yaml
helm list -n angelowl