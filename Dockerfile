FROM golang:1.15.6

WORKDIR /app

COPY . .

RUN go mod vendor
RUN go build -o ./bin/rest-kafka ./cmd/main.go

ENV PORT=8085 \
    KAFKA_SERVER=kafka:9092 \
    KAFKA_TOPIC=topic

CMD [ "./bin/rest-kafka" ]