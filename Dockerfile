# syntax=docker/dockerfile:1

FROM golang:1.19.3-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY src/ ./src
# COPY .env.example ./.env

# RUN set -a && . ./.env

# Empty env file to bypass the check since we're making use of Coolify secrets
RUN touch ./.env


RUN go build -o /intership-microservice src/*.go

EXPOSE 8080

CMD [ "/intership-microservice" ]
