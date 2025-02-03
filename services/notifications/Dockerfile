FROM golang:1.23

WORKDIR /app

# Copiar go.mod y go.sum desde la raíz del proyecto
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código del servicio notifications
COPY services/notifications/ ./services/notifications/

# Establecer el directorio de trabajo al del servicio notifications
WORKDIR /app/services/notifications

# Construir el binario del servicio notifications
RUN go build -o notifications-service

# Comando para ejecutar el servicio
EXPOSE 8000
CMD ["./notifications-service"]