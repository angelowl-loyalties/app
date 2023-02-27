helm repo add kafka-ui https://provectus.github.io/kafka-ui
helm install kafka-ui kafka-ui/kafka-ui -n kafka-ui --create-namespace -f values.yaml --debug --atomic