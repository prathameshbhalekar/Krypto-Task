# Krypto-Task
Task: [krypto-task.pdf](https://github.com/prathameshbhalekar/Krypto-Task/files/9265000/krypto-task.pdf)
## Instalation
### Fork

Create your own copy of the project on GitHub. You can do this by clicking the Fork button  on the top right corner of the landing page of the repository.

### Clone

Note: For this you need to install [git](https://git-scm.com/downloads) on your machine

```bash
$ git clone https://github.com/YOUR_GITHUB_USER_NAME/Krypto-Task
```
where YOUR_GITHUB_USER_NAME is your GitHub handle.

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
JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT= 43200
CACHE_TIMEOUT= 180000

EMAIL_ID= <YOUR_EMAIL_ID>
EMAIL_PASSWORD= <PASSWORD>
```

### Run Docker Compose
```bash
$ docker-compose up -d --build
```
Note: For this you need to install [docker](https://docs.docker.com/engine/install/) on your machine

### You are up and running on port 8080!

## Documentation
Postman docs link: https://documenter.getpostman.com/view/13627665/Uzdv28Bz

## Solution for alerts
- Subscribed to binance websocket using coroutines
- Published triggered alert to redis PUBSUB on every message of the socket
- Subscribed to the PUBSUB using coroutines and sent mails using net/smtp

## Task Checklist
- [x] Create, delete and fetch alert endpoints with filters and pagination for fetch
- [x] User auth with JWT
- [x] Binance websocket for triggers
- [x] Complete email functionality with SMTP
- [x] Caching for fetch alerts with redis
- [x] Redis PUBSUB for email message queue
- [x] Containerize with docker 
- [x] Work at Krypto ʕᵔᴥᵔʔ
