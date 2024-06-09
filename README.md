# Реализация REST API сервера на golang с использованием фреймворка gin
Сервер может сохранять данные из HTTP POST запросов и находить данные по id

### Установка
1. Скопировать репозиторий
```bash
git clone github.com/wowlikon/rest_db_question
cd rest_db_question
```

2. Установить зависимости
```bash
go mod tidy
```

3. Запуск
```bash
go build && rest_db_question
```

### Методы работы с API

#### POST `/address`
Позволяет отправить данные на сервер
```bash
curl -X POST --json '{"name": "alex", "address": "home", "longitude": 123, "latitude": 456}' http://localhost:8080/address

{"id": 1234567890}
```

#### GET `/address/:id`
Позволяет получить данные по id
```bash
curl -X GET http://localhost:8080/address/1234567890

{"name": "alex", "address": "home", "longitude": 123, "latitude": 456}
```
