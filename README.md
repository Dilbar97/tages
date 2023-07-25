# tages

## Содержание
* [Установка и настройка](#установка-и-настройка)
* [Миграция](#миграция)
* [Запуск](#запуск)

# Установка и настройка

```bash
$ git clone git@github.com:Dilbar97/tages.git

- в корне проекта найдите файл config.local.yaml
- скопируйте этот файл, вставьте там же и переименуйте копию на config.yaml

$ go mod tidy
$ docker-compose up -d --build
$ docker ps
  - скопируйте ID(хеш) postgres контейнера(скрин 1)
$ docker inspect {ID(хеш) postgres контейнера}
  - скопируйте поле IPAddress(скрин 2)
  - скопированный ip address вставьте в файле config.yaml поле db следующим образом: 
      db: postgres://tages:tanos@{IPAddress из скрин 2}:5432/media 
      например - db: postgres://tages:tanos@172.21.0.2:5432/media 
```
Скрин 1
![img.png](img.png)

Скрин 2
![img_1.png](img_1.png)

# Миграция

1) Следущим шагом нужно установить утилиту для запуска миграции

### MacOS
```bash
$ brew install golang-migrate
```

### Windows
```bash
$ scoop install migrate
```

### Linux (*.deb package)
```bash
$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
$ apt-get update
$ apt-get install -y migrate
```

2) Запуск самих миграции (скрин 3)

```bash
$ migrate -path db/migration -database "postgres://tages:tanos@{IPAddress из скрин 2}:5432/media?sslmode=disable" -verbose up
```

Скрин 3
![img_2.png](img_2.png)

# Запуск

### Server

```bash
$ go run main.go
```
Сервер сохранит данные и в диск и в базе (скрин 4)

Скрин 4
![img_4.png](img_4.png)

### Client

```bash
$ go run cmd/client/main.go {вторым аргументом передаётся путь к файлу, который хотите загрузить}

Пример: go run cmd/client/main.go /home/dilok/Изображения/test.jpg (скрин 5)
```

Скрин 5
![img_3.png](img_3.png)


## List запрос

![img_5.png](img_5.png)

## GetMedia запрос

![img_6.png](img_6.png)

