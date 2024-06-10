# Реализация REST API сервера на golang с использованием фреймворка gin
Сервер может сохранять данные из HTTP POST запросов и находить данные по id

### Установка
1. Скопировать репозиторий
```shell
git clone github.com/wowlikon/rest_db_question
cd rest_db_question
```

2. Установить зависимости
```shell
go mod tidy
```

2. Создать `.env`
```properties
admin_email=admin@gmail.com
bot_email=server_bot@mail.ru
bot_password=bot_password123
```

3. Запуск
```shell
go build && rest_db_question
```

### Методы работы с API

#### POST `/address`
Позволяет отправить данные на сервер
```shell
curl -X POST --json '{"name": "alex", "address": "home", "longitude": 123, "latitude": 456}' http://localhost:8080/address

{"id": 1234567890}
```

#### GET `/address/id/:id`
Позволяет получить данные по id
```shell
curl -X GET http://localhost:8080/address/1234567890

{"name": "alex", "address": "home", "longitude": 123, "latitude": 456}
```

#### GET `/address/name/:name`
Позволяет получить данные по name
```shell
curl -X GET http://localhost:8080/address/alex

{"name": "alex", "address": "home", "longitude": 123, "latitude": 456}
```
