# crypto-trade
[![Lint](https://github.com/mmontes11/crypto-trade/workflows/Lint/badge.svg)](https://github.com/mmontes11/crypto-trade/actions?query=workflow%3ALint)
[![Build](https://github.com/mmontes11/crypto-trade/workflows/Build/badge.svg)](https://github.com/mmontes11/crypto-trade/actions?query=workflow%3ABuild)
[![Test](https://github.com/mmontes11/crypto-trade/workflows/Test/badge.svg)](https://github.com/mmontes11/crypto-trade/actions?query=workflow%3ATest)
[![Release](https://github.com/mmontes11/crypto-trade/workflows/Release/badge.svg)](https://github.com/mmontes11/crypto-trade/actions?query=workflow%3ARelease)
[![Go Report Card](https://goreportcard.com/badge/github.com/mmontes11/crypto-trade)](https://goreportcard.com/report/github.com/mmontes11/crypto-trade)
[![Go Reference](https://pkg.go.dev/badge/github.com/mmontes11/crypto-trade.svg)](https://pkg.go.dev/github.com/mmontes11/crypto-trade)

Simulates real time cryptocurrency tradings in a microservice ecosystem.

---

- [Tech stack](#tech-stack)
- [How it works](#how-it-works)
- [Installation](#installation)
- [API](#api)
- [Production environment](#production-environment)

---

## Tech stack

- [Go](https://golang.org/)
- [Nats](https://nats.io/)
- [ClickHouse](https://clickhouse.tech/)
- [Docker](https://www.docker.com/)
- [Kubernetes](https://kubernetes.io/)
- [Helm](https://helm.sh/)
- [GitHub actions](https://github.com/features/actions)

## How it works

It contists in 3 microservices:
- **publisher**: Generates random trades and sends them to Nats.
- **subscriber**: Receives trades and stores them in ClickHouse.
- **api**: Exposes a REST API with the trades from ClickHouse.

## Installation

###### Local

1. Set up Nats and ClickHouse:
```bash
$ docker-compose -f docker-compose.services.yml up -d
```
2. Start each microservice in a different tab by manually running:
```bash
$ make run
```
###### Local + Tmux

Install [tmux](https://github.com/tmux/tmux) and then:
```bash
$ ./scripts/run-dev.sh
```
###### Docker

```bash
$ ./scripts/run-docker.sh
```
###### Kubernetes

```bash
$ ./scripts/deploy.sh
```
## API

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/155f5c6f3ba941caed61#?env%5Bcrypto-trade%20PRO%5D=W3sia2V5IjoidXJsIiwidmFsdWUiOiJodHRwczovL2NyeXB0by10cmFkZS5tbW9udGVzLWRldi5kdWNrZG5zLm9yZyIsImVuYWJsZWQiOnRydWV9XQ==)

###### GET /health

Indicate the service is up and running.

Responses:
- **200 OK**

###### GET /api/trades

Retrieves real time cryptocurrency trades. 

Query params:
- **groupBy**: second | minute | hour
- **crypto**: btc | eth
- **currency**: eur | usd
- **fromDate**: 1970-01-01T00:00:00Z
- **toDate**: 2050-01-01T00:00:00Z
- **limit**: 100

Responses:
- **200 OK**
- **404 Not Found**

Body:
- Array of trades


## Production environment

[https://crypto-trade.mmontes-dev.duckdns.org/api/trades?groupBy=hour&crypto=btc&currency=eur](https://crypto-trade.mmontes-dev.duckdns.org/api/trades?groupBy=hour&crypto=btc&currency=eur)
