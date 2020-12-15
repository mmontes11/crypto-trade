FROM golang:1.15-alpine3.12 AS builder

ARG SERVICE

ENV WORKDIR /go/src/crypto-trade
RUN mkdir -p ${WORKDIR}
WORKDIR ${WORKDIR}

COPY . .

RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o bin/app cmd/$SERVICE/main.go

FROM alpine:3.12.0

COPY --from=builder /go/src/crypto-trade/bin/app /app

CMD [ "/app" ]