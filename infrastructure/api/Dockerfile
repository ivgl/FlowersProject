# Этап сборки
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Копируем файлы проекта
COPY . .

# Скачиваем зависимости
RUN go mod download

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o plant-tracker ./cmd/server

# Финальный этап
FROM alpine:latest

WORKDIR /app

# Копируем бинарный файл из этапа сборки
COPY --from=builder /app/plant-tracker .

# Устанавливаем curl для healthcheck
RUN apk add --no-cache curl

# Запускаем приложение
CMD ["/bin/sh", "-c", "env && ./plant-tracker"] 