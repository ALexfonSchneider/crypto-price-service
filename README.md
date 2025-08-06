# Crypto Price Service

Сервис для отслеживания криптовалют.
Позволяет управлять watchlist, получать цены и следить за изменениями. 
Поддерживает интеграцию с внешними API (реализован CoinGecko) и автоматическое обновление цен.

---

## Функционал

- ✅ Добавление монеты в watchlist по символу (`BTC`, `ETH`)
- ✅ Удаление монеты
- ✅ Получение списка отслеживаемых монет
- ✅ Получение цен по монете и временной метке
- ✅ Поиск ближайшей цены к заданному времени
- ✅ Автоматические миграции БД
- ✅ REST API с документацией Swagger
- ✅ Кастомные ошибки с кодами и HTTP-статусами
- ✅ Логирование через `slog`
- ✅ Поддержка Docker и `docker-compose`
- ✅ Конфигурация через YAML
- ✅ Внешние клиенты (CoinGecko)

---

## Технологии

| Слой | Технология                                                 |
|------|------------------------------------------------------------|
| Язык | Go 1.23.1                                                  |
| Веб-фреймворк | [Gin](https://gin-gonic.com/)                              |
| База данных | PostgreSQL                                                 |
| ORM / SQL-генератор | [sqlc](https://sqlc.dev/)                                  |
| Миграции | [goose](https://github.com/pressly/goose)                  |
| Документация API | [Swaggo](https://github.com/swaggo/swag) (Swagger/OpenAPI) |
| DI | Ручная (чистые зависимости)                                |
| Конфигурация | koanf                                                      |
| Логирование | `log/slog`                                                 |

---

## Установка и запуск

### 1. Клонировать репозиторий

```bash
git clone https://github.com/ALexfonSchneider/crypto-price-service
cd crypto-price-service
```

### 2. Запустить через Docker Compose

```bash
docker-compose up --build
```

Сервис будет доступен на:

- **API**: `http://localhost:8080/api/v1`
- **Swagger UI**: `http://localhost:8080/swagger/index.html`

>  При запуске:
> - Применяются миграции
> - Сервис стартует

---

## Примеры запросов

### Добавить монету

```bash
curl -X POST http://localhost:8080/api/v1/coins \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Bitcoin",
    "symbol": "BTC"
  }'
```

### Удалить монету

```bash
curl -X DELETE http://localhost:8080/coins/BTC
```

### Получить список монет

```bash
curl http://localhost:8080/api/v1/coins
```

### Получить все цены для BTC

```bash
curl http://localhost:8080/prices/BTC
```

### Получить цену, ближайшую к определённому времени (в ms)

```bash
curl "http://localhost:8080/prices/BTC/closest?timestamp=1717000000000"
```

---

## 📁 Структура проекта

```
internal/
├── client/           # Внешние API (CoinGecko)
├── config/           # Конфигурация
├── db/               # SQLC-генерация, SQL-запросы
├── delivery/         # HTTP-хэндлеры, middleware и другой транспорт
├── dto/              # Data Transfer Objects
├── errors/           # Кастомные ошибки
├── models/           # Доменные модели
├── repository/       # Работа с БД (PostgreSQL)
├── services/         # Бизнес-логика
└── watcher/          # Сервис обновления цен монет
```

---

## Swagger UI

[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

Чтобы перегенерировать документацию:

```bash
make swagger
```

---

## Конфигурация

Конфигурация находится в `config/`:

- `local.yaml` — для локальной разработки
- `docker.yaml` — для Docker

Выбрать конфиг можно с помощью переменной окружения **APP_ENV**
Просто запустить:
```bash
APP_ENV=local go run cmd/service/main.go
```

Поддерживает:
- Порт HTTP
- Настройки PostgreSQL
- Настройки внешних API

---

## 🧹 Миграции

Миграции находятся в папке `migrations/` и применяются посредством:
```bash
 go run cmd/migrate/main.go
```

### Добавить новую миграцию

```bash
migrate create -ext sql -dir migrations -seq add_price_table
```

## Для разработки

### Установка зависимостей

```bash
go mod download
```

### Установка инструментов (один раз)

```bash
make swagger-install
```
```bash
make sqlc-install
```

### Запуск без Docker

1. Запустите PostgreSQL
2. Примените миграции:

```bash
make migrate
```

3. Запустите сервис:

```bash
make run
```

## Docker

Образ собирается с помощью `Dockerfile`:

Собрать:

```bash
docker build -t crypto-price-service .
```