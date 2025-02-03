FROM golang:1.23

WORKDIR /app

# Copiar go.mod y go.sum desde la raíz del proyecto
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código del servicio inventory
COPY services/inventory/ ./services/inventory/

# Establecer el directorio de trabajo al del servicio inventory
WORKDIR /app/services/inventory

# Construir el binario del servicio inventory
RUN go build -o inventory-service

# Comando para ejecutar el servicio
EXPOSE 8000
CMD ["./inventory-service"]