FROM golang:1.25.4-alpine3.22 AS builder

WORKDIR /app

COPY ../../go.mod ./
COPY ../../go.sum ./

RUN go mod download

COPY ../../ ./

WORKDIR /app/cmd/api
RUN go build -o /mealbook

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /mealbook .

EXPOSE 8080
CMD ["./mealbook"]
