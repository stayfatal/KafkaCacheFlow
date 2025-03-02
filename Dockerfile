FROM golang:1.22-alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

FROM builder AS gateway_builder

WORKDIR /app

COPY ./domain ./

COPY ./services/gateway ./services/gateway

RUN go build -o app ./services/gateway/cmd/main.go

FROM alpine:3.20 AS gateway

COPY --from=gateway_builder /app /app

CMD [ "./app/app" ]

FROM builder AS worker_builder

WORKDIR /app

COPY ./domain ./

COPY ./services/worker ./services/worker

RUN go build -o app ./services/worker/cmd/main.go

FROM alpine:3.20 AS worker

COPY --from=worker_builder /app /app

CMD [ "./app/app" ]