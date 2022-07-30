# Krypto-Task

## Instalation

### Setup the .ENV file
Add JWT secret key, your email and password

```bash
DB_USER= user 
DB_PASS= password
DB_HOST= database
DB_NAME= user

RDB_ADDR= redis:7000
RDB_PASS= password123

JWT_SECRET_KEY= <JWT_SECRET_KEY>
JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT= 6000000
CACHE_TIMEOUT= 10000000

EMAIL_ID= <YOUR_EMAIL_ID>
EMAIL_PASSWORD= <PASSWORD>
```

### Run Docker Compose
```bash
$ docker-compose up -d --build
```

### You are up and running on port 8080!

## Documentation
Postman docs link: https://documenter.getpostman.com/view/13627665/Uzdv28Bz

## Solution for alerts
- Subscribed to binance websocket using coroutines
- Published triggered alert to redis PUBSUB on every message of the socket
- Subscribed to the PUBSUB using coroutines and sent mails using net/smtp
