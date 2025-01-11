FROM golang:1.20

WORKDIR /app

# Copiar y descargar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . ./

# Compilar la aplicación
RUN go build -o inventory

# Ejecutar la aplicación
CMD ["./inventory"]
