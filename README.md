# image-processing-backend
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)


## 📑 Содержание

- [Описание](#-описание)
- [Используемые технологии](#-используемые-технологии)
- [Запуск](#-запуск)

## 📖 Описание

Этот проект представляет собой бэкенд-сервис на Go, предназначеный для обработки изображений посредством различных фильтров.
Функциональность - регистрация/авторизация, загрузка и хранение изображений, получение статуса задачи и изменённого изображения.

## 🛠️ Используемые технологии
- PostgreSQL (хранение User и Task)
- Redis (кэширование сессий)
- Swagger (документация API)
- Docker и docker compose (развёртывание сервиса)
- CI (GitHub actions)

## ⚙️ Запуск

```bash
git clone https://github.com/kasparovgs/image-processing-backend.git
cd image-processing-backend
make launch_services
```


Посмотреть swagger-документацию можно здесь:
[http://localhost:8080/swagger/docs/index.html](http://localhost:8080/swagger/docs/index.html)

RabbitMQ management: [http://localhost:15672/#/](http://localhost:15672/#/) (login: guest, password: guest)

Рекомендуемый сервис для конвертации изображений в base64 и обратно: [https://imagestoolkit-develop.netlify.app/base64](https://imagestoolkit-develop.netlify.app/base64)
