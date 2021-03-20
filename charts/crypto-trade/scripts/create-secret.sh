#!/usr/bin/env bash

set -euo pipefail

# export CLICKHOUSE_URL='tcp://crypto-trade-clickhouse:9000?username=default&password=default&database=default'
# export NATS_URL='nats://crypto-trade-nats:422'
# export PERSONAL_ACCESS_TOKEN='PERSONAL_ACCESS_TOKEN'

kubectl create secret generic crypto-trade \
  --from-literal=CLICKHOUSE_URL='$CLICKHOUSE_URL' \
  --from-literal=NATS_URL='$NATS_URL' \
  --from-literal=MIGRATIONS_URL='github://mmontes11:$PERSONAL_ACCESS_TOKEN@mmontes11/crypto-trade/cmd/subscriber/model/migrations'
