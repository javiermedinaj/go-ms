FROM golang:1.23

WORKDIR /app

# Copiar go.mod y go.sum desde la raíz del proyecto
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código del servicio orders
COPY services/orders/ ./services/orders/

# Establecer el directorio de trabajo al del servicio orders
WORKDIR /app/services/orders

# Construir el binario del servicio orders
RUN go build -o orders-service

# Comando para ejecutar el servicio
EXPOSE 8000
CMD ["./orders-service"]