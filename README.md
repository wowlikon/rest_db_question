# Реализация REST API сервера на golang с использованием фреймворка gin
Сервер может сохранять данные из HTTP POST запросов и находить данные по id

### POST `/address`
```bash
curl -X POST --json '{"name": "alex", "address": "home", "longitude": 123, "latitude": 456}' http://localhost:8080/address

{"id": 1234567890}
```

### GET `/address`
```bash
curl -X GEt --json '{"id": 1234567890}' http://localhost:8080/address

{"name": "alex", "address": "home", "longitude": 123, "latitude": 456}
```

# ⚠️Сервер ещё на этапе разработки и может работать не стабильно/не работать⚠️
