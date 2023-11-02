# Etapa de compilación para el servicio de archivos
FROM golang:1.21.1-alpine AS build-smb-rest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux go build -o smb-rest ./cmd/rest

# Imagen final para el servicio de archivos
FROM alpine:latest

WORKDIR /app

COPY --from=build-smb-rest /app/smb-rest .

EXPOSE 8010

CMD ["./smb-rest"]