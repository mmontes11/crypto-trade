#!/usr/bin/env bash

RELEASE="crypto-trade"
REPO="mmontes"
NAMESPACE="crypto-trade"

helm repo add "$REPO" https://charts.mmontes-dev.duckdns.org
helm repo update

echo "🚀 Deploying '${RELEASE}'..."
helm upgrade --install "$RELEASE" "$REPO/$RELEASE" --namespace "$NAMESPACE"
