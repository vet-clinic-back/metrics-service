# Используем официальный образ Go
FROM golang:1.19-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем исходный код
COPY . .

# Собираем приложение
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/main.go

# Финальный образ
FROM alpine:3.18

WORKDIR /app

# Копируем бинарный файл из builder
COPY --from=builder /app/main /app/main

# Копируем конфиги
COPY config.yaml /app/config.yaml

# Открываем порты
EXPOSE 8080 8081

# Запускаем приложение
CMD ["/app/main"]