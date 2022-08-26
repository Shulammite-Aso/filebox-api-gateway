# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS build

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /fb-api-gateway

FROM scratch

WORKDIR /

COPY --from=build /fb-api-gateway /fb-api-gateway

EXPOSE 3000

CMD [ "/fb-api-gateway" ]