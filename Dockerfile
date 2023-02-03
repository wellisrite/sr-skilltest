FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

ENV SQL_DATABASE=$SQL_DATABASE


RUN psql -U postgres -d ${SQL_DATABASE} -f /internal/infra/db/base_migration.sql

RUN go mod tidy

RUN go mod vendor

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./sr-skilltest"]
