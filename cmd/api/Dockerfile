FROM golang:1.15-alpine3.12 AS builder

ENV WORKDIR /go/src/api
RUN mkdir -p ${WORKDIR}
WORKDIR ${WORKDIR}

COPY . .

RUN go build -o bin/api -v .

FROM alpine:3.12.0

COPY --from=builder /go/src/api/bin/api /app/api

CMD [ "/app/api" ]