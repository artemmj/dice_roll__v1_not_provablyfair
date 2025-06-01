## dice_roll - gRPC-сервис-игра "Подбрасывание кубика"

#### Установка
1. Скопировать репозиторий, перейти в папку
```
git clone git@github.com:artemmj/dice_roll__v1_not_provablyfair.git

cd dice_roll__v1_not_provablyfair/
```
2. Создать файл с переменными окружения .env для докера (есть пример .env_example). При первом запуске создастся новая БД, применятся миграции. В файле уканазы дефолтные параметры подключения, при необходимости можно поменять.
```
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=postgres
```
4. Выполнить команду ```docker-compose up -d --build``` чтобы собрать контейнеры
5. В папке volumes/ появятся данные БД
6. Чтобы сыграть, используем например Postman, подключимся к localhost:50051 by gRPC и выполним Play:
7. Чтобы исполнить тесты, выполните команду при работающем контейнере
```
go test ./tests -count=1 -v
```
