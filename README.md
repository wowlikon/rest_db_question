# Реализация REST API сервера на golang с использованием фреймворка gin
Сервер может сохранять данные из HTTP POST запросов и находить данные по id

### POST `/address`
Позволяет отправить данные на сервер
```bash
curl -X POST --json '{"name": "alex", "address": "home", "longitude": 123, "latitude": 456}' http://localhost:8080/address

{"id": 1234567890}
```

### GET `/address/:id`
Позволяет получить данные по id
```bash
curl -X GET http://localhost:8080/address/1234567890

{"name": "alex", "address": "home", "longitude": 123, "latitude": 456}
```

### GET `/all`
Позволяет получить все данные
```bash
curl -X GET http://localhost:8080/all

{"1234567890":{"name":"alex","address":"home","longitude":123,"latitude":456}}
```