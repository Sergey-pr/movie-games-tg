English documentation is <a href=https://github.com/Sergey-pr/movie-games-tg/tree/main#moviegames-telegram-miniapp>here</a>

# КИНОИГРЫ Telegram MiniApp

КИНОИГРЫ это приложение для телеграма. Оно представляет собой карточную игру в которой 
нужно угадать фильм по подсказкам. В неё интересно поиграть и вы можете использовать её 
как основу, чтобы создать своё приложение.

<div align="center">
    <img src=https://github.com/Sergey-pr/movie-games-tg/raw/main/assets/preview.gif
    alt="Game preview">
</div>

# Структура проекта

* Бэкенд построен на <a href=https://go.dev/>Go</a> с помощью 
  <a href=https://github.com/gorilla/mux#gorillamux>Gorilla Mux</a> 
  http сервера и <a href=https://github.com/doug-martin/goqu#readme>goqu</a>
  для работы с базой
* Фронтенд это приложение на <a href=https://vuejs.org/>Vue.js</a>
  фреймворке с <a href=https://router.vuejs.org/>Vue Router</a> и
  <a href=https://vuex.vuejs.org/>Vuex</a>.
* База данных это <a href=https://www.postgresql.org/>PostgreSQL</a> с 
  <a href=https://github.com/amacneil/dbmate#dbmate>Dbmate</a> для миграций

## Backend

Бэкенд разделён на несколько модулей:
* **Models** - Модели объектов для таблиц базы данных. У каждой модели есть методы для 
  создания, обновления или удаления модели. Так же там есть функции для получения
  объектов из базы данных.
* **Forms** - Объекты в которых хранятся входные данные на бэкенд. Они используются, чтобы
  обработать тело запроса в объект и потом уже работать с его полями.
* **Serializers** - Объекты в которых хранятся выходные данные на фронтенд. Они
  используются, чтобы разделить объекты базы данных с объектами ответов API.
* **Handlers** - Функции которые обрабатывают API запросы.

Так же на бэкенде есть следующие папки:
* **bot_files** - Здесь мы храним файлы, которые используются ботом для ответов в чат.
* **card_files** - Здесь мы храним изображения карточек. В хендлерах есть хендлер который
  отдаёт эти файлы по их id. Каждый файл назван по id.
* **config** - Здесь обрабатываются переменные окружения.
* **migrations** - Это папка для миграций
  <a href=https://github.com/amacneil/dbmate#dbmate>Dbmate</a>.
* **persist** - Здесь хранятся подключения которые активны пока запущено приложение.
  Например, соединение с базой данных.
* **utils** - Здесь мы храним утилиты и функции хелперы.

Среди моделей есть специальный объект с названием **CardProcessor**.
Это объект с помощью которого можно поэтапно добавить новые карточки через чат с ботом.
Там обрабатывается 16 этапов для каждого из полей карточек. Некоторые поля дублируются
для разных языков, у них есть суффиксы `en` или `ru`.

## Frontend

Фронтенд состоит из 3 папок

* **assets** - Статичные файлы
* **components** - Компоненты Vue
* **services** - Утилиты и функции помощники

Всего 4 основных компонента и несколько дополнительных:

* **LandingPage** - Основная страница проекта. Отсюда можно запустить игру,
  открыть страницу с правилами или страницу с лучшими игроками
* **PlayGame** - Страница самой игры, она подгружает дополнительный компонент в 
  зависимости от этапа в котором находится игра:
1. Этап `play` - Показывает компонент **CardPage** с вопросами из карточки
2. Этап `info` - Показывает компонент **CardInfoPage** с информацией о фильме карточки
3. Этап `end` - Показывает компонент **TheEndPage** в котором указано сколько очков вы заработали за игру
* **RulesPage** - Страница с правилами игры
* **LeaderboardPage** - Страница со списком лучших игроков

Так же есть небольшой компонент **LoadingComponent** - это svg анимация лоадера для страниц

В папке **services** 2 файла:
* **api.js** - Объект axios для работы с запросами
* **utils.js** - Функции хелперы

## База данных

База данных состоит из 6 таблиц

* **answers** - Хранит в себе ответы пользователей на карточки, чтобы потом выводить
  их в таблице лучших игроков.
* **bot_files** - Сохраняет id файлов которые бот отправляет в чаты телеграм,
  чтобы не загружать их заново каждый раз.
* **card_processors** - Сохраняет в себе этап создания новой карточки для каждого пользователя.
* **cards** - Хранит в себе данные по карточкам. Многие поля дублируются для переводов
  у таких ролей будут суффиксы `ru` и `en`
* **schema_migrations** - Сохраняет примененные миграции
  <a href=https://github.com/amacneil/dbmate#dbmate>Dbmate</a>
* **users** - Хранит информацию о пользователях их id, username, язык и т.д.

# Схема работы приложения

Приложение работает на 2 языках Английском и Русском в зависимости от языка в telegram пользователя.
Так же можно изменить язык на главной странице приложения.

1. При запуске бэкенда отправляется запрос на регистрацию колбеков бота на хендлер `bot-updates`.
2. Если написать в чат бота `/start` или любое сообщение, бот будет отвечать на них через метод
   `POST /api/public/bot-updates/`
3. При загрузке фронтенда мы получаем `initData` от телеграма и отправляем её на бэкенд `POST /api/public/login/`
   где проверяем её по хэшу, и если всё ок, то отправляем JWT токен в ответе.
   На фронтенде токен сохраняется и используется для отправки запросов на приватные методы.
4. Сверху приложения есть переключатель языка. При смене языка мы сохраняем выбранный
   язык пользователя с помощью метода `POST /api/user/lang/`
5. Когда нажимаем кнопку играть, м отправляемся на страницу PlayGame, где фронтенд
   получает все карточки с помощью метода `GET /api/cards/`, перемешивает их и отдаёт по очереди.
6. Изображения карточек сохранены как id и чтобы получить их мы используем метод
   `GET /api/public/bot-image/{image_id}/`
7. После того как пользователь ответил на карточку мы отправляем запрос на сохранение 
   ответа и количества очков с помощью метода `POST /api/user/answer/` и открываем страницу с 
   информацией по фильму.
8. Когда пользователь ответил на все карточки мы открываем страницу с общим количеством 
   заработанных очков. Оттуда можно вернуться на главную страницу.
9. С главной страницы можно открыть страницу правил или список лучших игроков. Лучших
   игроков получаем методом `GET /api/leaderboard/`.
10. Если вы установите пользователя админом (в базе данных в таблице `users` столбце `is_admin` поставите `true`)
    то вы сможете воспользоваться командой `/add` в чате бота, после которой запустится алгоритм добавления новой 
    карточки. Бот подскажет по этапно как добавить новую карточку.

# Установка и запуск

## Используя Docker compose

Можно запустить приложение через 
<a href=https://github.com/Sergey-pr/movie-games-tg/blob/main/docker-compose.yml>
docker-compose.yml</a> файл.

<details><summary><b>Показать инструкцию</b></summary>

1. Установите <a href=https://www.docker.com/>Docker</a> и убедитесь что с ним установился
   <a href=https://docs.docker.com/compose/gettingstarted/>Docker Compose</a>.
2. Сделайте копию файла 
   <a href=https://github.com/Sergey-pr/movie-games-tg/blob/main/docker-compose.yml>
   docker-compose.yml</a> и назовите её docker-compose.yml.local
3. Измените значения **Переменных Окружения** на нужные вам

```dotenv
# frontend

# VUE_APP_BASE_URL это адрес бэкенда, туда будут идти запросы с фронтенда
VUE_APP_BASE_URL=localhost:8888


#backend

# DATABASE это DSN строка подключения к базе данных
DATABASE='user=postgres password=postgres host=localhost port=5432 dbname=movie_games sslmode=disable'
# JWT_TOKEN это строка секрет для генерации JWT токенов
JWT_TOKEN=sfhjahkfg8749GHGJHgjhds
# TELEGRAM_BOT_TOKEN это токен вашего телеграм бота, его можно получить у @BotFather 
# в телеграм при регистрации бота
TELEGRAM_BOT_TOKEN=123456789:qwertyuioASDFGHJKLzxcvbnm
# FRONTEND_HOSTNAME это адрес фронтенда
FRONTEND_HOSTNAME=localhost:8080
# BACKEND_HOSTNAME это адрес бэкенда
BACKEND_HOSTNAME=localhost:8888


# migrator
# Это сервис применяющий миграции

# DBMATE_NO_DUMP_SCHEMA это переменная для создания схемы базы. При true файл схемы не создаётся
# т.к. он нам не нужен тут стоит true
DBMATE_NO_DUMP_SCHEMA=true
# DBMATE_MIGRATIONS_DIR это папка с миграциями Dbmate у нас это migrations
DBMATE_MIGRATIONS_DIR=migrations
# DATABASE_URL это строка подключения к базе, теперь в другом виде
DATABASE_URL=postgres://postgres:postgres@db:5432/movie_games?sslmode=disable


# database
# Здесь указываются логопасы базы
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=movie_games
```
4. Измените файл `nginx.conf` на нужные вам настройки nginx
5. Запустите команду `docker-compose -f docker-compose.yml.local build` чтобы собрать докер контейнеры
6. Запустите `docker-compose -f docker-compose.yml.local up -d` чтобы запустить проект
7. После этого колбеки от сообщений телеграма привяжутся к адресу бота, но будет нужно еще
   установить адрес фронтенда на кнопку боту у @BotFather чтобы приложение можно было открыть по кнопке.

</details>

## Запустить всё отдельно

Для разработки или отладки советую запускать все модули отдельно.

### База данных

<details><summary><b>Показать инструкцию</b></summary>

1. Создайте базу данных <a href=https://www.postgresql.org/>PostgreSQL</a>
2. Установите <a href=https://github.com/amacneil/dbmate#dbmate>Dbmate</a>
3. В папке backend заполните .env файл с **Переменными Окружения** для 
   <a href=https://github.com/amacneil/dbmate#dbmate>Dbmate</a>
```dotenv
# DBMATE_NO_DUMP_SCHEMA это переменная для создания схемы базы. При true файл схемы не создаётся
# т.к. он нам не нужен тут стоит true
DBMATE_NO_DUMP_SCHEMA=true
# DBMATE_MIGRATIONS_DIR это папка с миграциями Dbmate у нас это migrations
DBMATE_MIGRATIONS_DIR=migrations
# DATABASE_URL это строка подключения к базе, теперь в другом виде
DATABASE_URL=postgres://postgres:postgres@db:5432/movie_games?sslmode=disable
```
4. Запустите команду `dbmate up` чтобы применить миграции
</details>

### Backend

<details><summary><b>Показать инструкцию</b></summary>

1. Установите <a href=https://go.dev/>Go</a> версии 1.20 или новее
2. Перейдите в папку backend
3. Укажите **Переменные Окружения**
```dotenv
# REST_LISTEN это адрес на котором бэкенд слушает запросы
REST_LISTEN=0.0.0.0:8888
# DATABASE это DSN строка подключения к базе данных
DATABASE='user=postgres password=postgres host=localhost port=5432 dbname=movie_games sslmode=disable'
# JWT_TOKEN это строка секрет для генерации JWT токенов
JWT_TOKEN=sfhjahkfg8749GHGJHgjhds
# TELEGRAM_BOT_TOKEN это токен вашего телеграм бота, его можно получить у @BotFather 
# в телеграм при регистрации бота
TELEGRAM_BOT_TOKEN=123456789:qwertyuioASDFGHJKLzxcvbnm
# FRONTEND_HOSTNAME это адрес фронтенда
FRONTEND_HOSTNAME=localhost:8080
# BACKEND_HOSTNAME это адрес бэкенда
BACKEND_HOSTNAME=localhost:8888
```
4. Запустите `go mod download` чтобы установить зависимости
5. Запустите `go build main.go` чтобы стартануть веб сервер бэкенда

</details>

### Frontend

<details><summary><b>Показать инструкцию</b></summary>

1. Установите <a href=https://nodejs.org/en>Node.js</a>
2. Перейдите в папку frontend
3. Запустите `npm install` чтобы установить зависимости
4. Запустите `npm run serve` чтобы запустить сервер фронтенда на localhost.
   Так же можно использовать `npm run dev` который сгенерирует статичные html
   файлы в папке dist и генерировать новые при каждом изменении. Их можно потом
   отдавать любым веб сервером, например:
   <a href=https://www.npmjs.com/package/serve>serve</a>.

</details>

### Запуск через Telegram

Тут есть 2 варианта:
* Создать аккаунт тестового окружения telegram и смотреть приложения с localhost 
  через `npm run serve`
* Сделать свои порты публичными через утилиту такую как
  <a href=https://ngrok.com/>ngrok</a> и тестировать приложение через обычный телеграм 
  аккаунт

Больше информации по второму методу:

<details><summary><b>Показать инструкцию</b></summary>

1. Зарегистрируйтесь на <a href=https://ngrok.com/>ngrok</a>
2. Настройте конфиг ngrok для работы с двумя портами
```yml
version: "2"
authtoken: yourNgrokToken
tunnels:
  backend:
    proto: http
    addr: 8888
  frontend:
    proto: http
    addr: 8080
```
3. Замените **Переменные Окружения** бэкенда на адреса из ngrok:
```dotenv
FRONTEND_HOSTNAME=https://2acf-188-233-88-176.ngrok-free.app
BACKEND_HOSTNAME=https://7cc1-188-233-88-176.ngrok-free.app
```
4. Запустите бэкенд с помощью `go build main.go`. Сервер запустится на порте 8888,
   порт можно изменить с помощью **Переменной Окружения** `REST_LISTEN`
5. Замените **Переменные Окружения** фронтенда на адрес из ngrok:
```dotenv
# VUE_APP_BASE_URL это адрес бэкенда, туда будут идти запросы с фронтенда
VUE_APP_BASE_URL=https://7cc1-188-233-88-176.ngrok-free.app
```
6. Запустите фронтенд командой `npm run dev` чтобы сгенерировалась папка dist
7. Запустите веб сервер <a href=https://www.npmjs.com/package/serve>serve</a>
   командой `serve -l 8080` в папке frontend
8. Укажите адрес фронтенда ngrok как адрес кнопки меню бота у @BotFather в телеграм
9. Теперь можете открыть приложение через telegram. Больше информации по отладке приложений в телеграм 
   <a href=https://core.telegram.org/bots/webapps#testing-mini-apps>Здесь</a>
</details>
