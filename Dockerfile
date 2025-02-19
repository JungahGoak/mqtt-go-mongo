FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o mqtt-consumer ./cmd

# GLIBC 2.34 포함된 걸로 변경함
FROM debian:bookworm

WORKDIR /app

COPY --from=builder /app/mqtt-consumer .

ENV MQTT_BROKER="mosquitto:1883"
ENV MONGO_URI="mongodb://mongo:27017/mqtt_data"

CMD ["./mqtt-consumer"]

