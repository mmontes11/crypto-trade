#!/bin/bash

set -e

kubectl apply -f manifests/namespace.yml

# Nats
echo "🚀    Deploying Nats..."
helm repo add nats https://nats-io.github.io/k8s/helm/charts/
helm upgrade nats nats/nats -f manifests/nats.values.yml -n crypto-trade --install

# ClickHouse
echo "🚀    Deploying Clickhouse..."
helm repo add liwenhe https://liwenhe1993.github.io/charts/
helm upgrade clickhouse liwenhe/clickhouse -f manifests/clickhouse.values.yml -n crypto-trade --install

# Microservices
for ms in $(ls -d cmd/*); do
    name=$(basename "$ms")
    manifests="$ms/manifests"
    echo "🚀    Deploying $name..."
    kubectl apply -f "$manifests"
done
