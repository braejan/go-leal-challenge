# Etapa de compilación para el servicio de archivos
FROM golang:1.21.3-alpine AS build-leal-challenge-rest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux go build -o leal-challenge-rest ./cmd/web/main.go

# Imagen final para el servicio de archivos
FROM alpine:latest

WORKDIR /app

COPY --from=build-leal-challenge-rest /app/leal-challenge-rest .

EXPOSE 8010

CMD ["./leal-challenge-rest"]