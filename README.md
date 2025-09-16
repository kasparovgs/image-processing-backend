# image-processing-backend
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)


## 📑 Содержание

- [Описание](#-описание)
- [Используемые технологии](#-используемые-технологии)
- [Запуск](#-запуск)
- [Полезные ссылки](#-полезные-ссылки)

## 📖 Описание

Этот проект представляет собой бэкенд-сервис на Go, предназначеный для обработки изображений посредством различных фильтров.
Функциональность - регистрация/авторизация, загрузка и хранение изображений, получение статуса задачи и изменённого изображения.

## 🛠️ Используемые технологии
- **PostgreSQL** (хранение User и Task)
- **Redis** (кэширование сессий)
- **Swagger** (документация API)
- **Docker/docker compose** (развёртывание сервиса)
- **CI** (GitHub actions)
- **Prometheus** (сбор и хранение метрик)
- **Grafana** (визуализация метрик)

## ⚙️ Команды

### Запуск
```bash
git clone https://github.com/kasparovgs/image-processing-backend.git
cd image-processing-backend
make launch_services
```

### Остановка
```bash
make stop_services
```

### Запуск тестов
```bash
make launch_with_tests
```

## 🔗 Полезные ссылки

| Инструмент / сервис | Ссылка | Примечание |
| --- | --- | --- |
| 📜 **Swagger UI** | [http://localhost:8080/swagger/docs/index.html](http://localhost:8080/swagger/docs/index.html) | Просмотр и тестирование API |
| 📨 **RabbitMQ Management** | [http://localhost:15672/#/](http://localhost:15672/#/) | Логин: `guest`, пароль: `guest` |
| 🖼 **Base64 converter** | [https://imagestoolkit-develop.netlify.app/base64](https://imagestoolkit-develop.netlify.app/base64) | Быстрая конвертация изображений в base64 и обратно |
| 📈 **Prometheus** | [http://localhost:9090/query](http://localhost:9090/query) | Сбор метрик |
| 📊 **Grafana** | [http://localhost:3000/login](http://localhost:3000/login) | Логин: `admin`, пароль: `admin` |
