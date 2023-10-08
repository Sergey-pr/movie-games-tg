Инструкция на русском языке <a href=https://github.com/Sergey-pr/movie-games-tg/blob/main/README_RU.md#%D0%BA%D0%B8%D0%BD%D0%BE%D0%B8%D0%B3%D1%80%D1%8B-telegram-miniapp>тут</a>

# MovieGames Telegram MiniApp

MovieGames is a Telegram Mini App. It's a card game where your
main objective is to guess movies by clues which are given to
you. It's fun to play, and you can use its source code to
create your own Mini Apps.

<div align="center">
    <img src=https://github.com/Sergey-pr/movie-games-tg/raw/main/assets/preview.gif
    alt="Game preview">
</div>

# Project Structure

* Backend is a web server written in <a href=https://go.dev/>Go</a> using 
  <a href=https://github.com/gorilla/mux#gorillamux>Gorilla Mux</a>
  for http server and <a href=https://github.com/doug-martin/goqu#readme>goqu</a>
  for database requests
* Frontend is a **Single Page Application** made on a <a href=https://vuejs.org/>Vue.js</a>
  framework using <a href=https://router.vuejs.org/>Vue Router</a> for routing and
  <a href=https://vuex.vuejs.org/>Vuex</a> for state management.
* Database is <a href=https://www.postgresql.org/>PostgreSQL</a> with 
  <a href=https://github.com/amacneil/dbmate#dbmate>Dbmate</a> for migrations

## Backend

Backend is separated to modules:
* **Models** - Object Related Mapping to database tables. Each model has methods for
  creating, updating and deleting an object. There are also functions to get object from
  database.
* **Forms** - Objects which represent data from requests to backend. We usually parse 
  request data to form, and then work with form object.
* **Serializers** - Objects which represent response data from backend. We usually get
  model from database and then serialize it before sending back to frontend or
  telegram bot api.
* **Handlers** - Functions to process requests to backend.

You can also notice next folders in the backend:
* **bot_files** - Here we have files that are used for bot responds.
* **card_files** - Here we save cards images to use later on the frontend.
  You can find a handler which serves this images in handlers module.
* **config** - Here we process Environment variables to use in the app.
* **migrations** - Is a folder with
  <a href=https://github.com/amacneil/dbmate#dbmate>Dbmate</a> migrations.
* **persist** - Here we have connections that persist while app is running.
  Right now it's only a database connection.
* **utils** - Here store utils and helper functions.

In the models there are special model which is called **CardProcessor**.
It's a state machine which lets you add new cards in the bot chat.
There are 16 states one for each of the cards fields. Some fields are duplicating 
for translations, these fields have affixes `en` or `ru`.

## Frontend

Frontend consist of 3 main folders

* **assets** - Static assets
* **components** - Vue components
* **services** - Utils and helpers

There are 4 main components:

* **LandingPage** - is a main page of the app, from here you can start the game,
  open rules page or leaderboard page
* **PlayGame** - is a main game component it loads a subcomponent depending on one of 3 states:
1. State `play` - is showing **CardPage** component with questions
2. State `info` - is showing **CardInfoPage** component which shows card movie info
3. State `end` - is showing **TheEndPage** component which tells you how many points you've got
* **RulesPage** - is showing page with hints how to play this game
* **LeaderboardPage** - is a page showing players with mot points

There are also a small **LoadingComponent** which is just a loader to use while app is loading

Services are 2 files:
* **api.js** - is an axios api instance with all the requests
* **utils.js** - contains helper functions

## Database

Database consist of 6 tables

* **answers** - saves how many points each user got on each card for teh first time
  for the leaderboard.
* **bot_files** - saves telegram file ids, so we don't upload files for every response
  and just use already uploaded files.
* **card_processors** - saves state for card processor for user
* **cards** - saves cards data some fields are duplicating with different affixes `ru` and `en`
  for translations.
* **schema_migrations** - saves applied migrations to database, it is a
  <a href=https://github.com/amacneil/dbmate#dbmate>Dbmate</a> technical table
* **users** - saves users data like username or preferred language

# App Workflow

App works in 2 languages English and Russian depending on user's telegram language.
You can also change language on app's main page.

1. On backend load we send request to register callbacks to backend `bot-updates` handler.
2. When typing `/start` to the bot, bot will respond with a welcome message using callback endpoint
   `POST /api/public/bot-updates/`
3. On app load we get telegram `initData` and send it to backend `POST /api/public/login/`
   where we validate it by hash, and if it is valid, we send JWT token in response
   to confirm that requests go from telegram users. Frontend stores JWT token and
   uses it for all private api requests.
4. On top of the main page there is a language switch. On switch we send request to save
   user's preferred language `POST /api/user/lang/`
5. When we press the play button, we go the PlayGame page, where frontend sends request
   to get all the cards `GET /api/cards/` then shuffle it and shows the first card.
6. Cards images are stored by ids, and we serve them from backend with
   `GET /api/public/bot-image/{image_id}/`
7. After user answered right or wrong we send request to save how many points user
   have got `POST /api/user/answer/` and show card info page.
8. When user answered on all the cards we show how many points user have got. From there
   user can go back to main page.
9. From main page user can go to the rules page which is static or leaderboard. 
   Leaderboard info we get with a `GET /api/leaderboard/` method.
10. If you set yourself as admin (database table `users` column `is_admin` set `true`) you
    can use bot command `/add` which will guide you through process of adding new cards
    to the game.
# Usage

## Docker compose

You can run this app with
<a href=https://github.com/Sergey-pr/movie-games-tg/blob/main/docker-compose.yml>
docker-compose.yml</a> file.

<details><summary><b>Show instructions</b></summary>

1. Setup <a href=https://www.docker.com/>Docker</a> and make sure you've got 
   <a href=https://docs.docker.com/compose/gettingstarted/>Docker Compose</a> with it.
2. Make a copy of
   <a href=https://github.com/Sergey-pr/movie-games-tg/blob/main/docker-compose.yml>
   docker-compose.yml</a> and name it docker-compose.yml.local
3. Edit **Environment Variables** to suit your needs
```dotenv
# frontend

# VUE_APP_BASE_URL is your backend address, requests to backend will go there
VUE_APP_BASE_URL=localhost:8888


#backend

# DATABASE is your database DSN string
DATABASE='user=postgres password=postgres host=localhost port=5432 dbname=movie_games sslmode=disable'
# JWT_TOKEN is your JWT token secret you can write here any combination of symbols
JWT_TOKEN=sfhjahkfg8749GHGJHgjhds
# TELEGRAM_BOT_TOKEN is your Telegram bot token you can get it from @BotFather 
# when registering your telegram bot
TELEGRAM_BOT_TOKEN=123456789:qwertyuioASDFGHJKLzxcvbnm
# FRONTEND_HOSTNAME is your frontend address
FRONTEND_HOSTNAME=localhost:8080
# BACKEND_HOSTNAME is your backend address
BACKEND_HOSTNAME=localhost:8888


# migrator
# This is a migrator service which will create all the tables in your database

# DBMATE_NO_DUMP_SCHEMA is boolean for createing dump schema file you can leave it 
# to true, as dump schema is not needed for this project
DBMATE_NO_DUMP_SCHEMA=true
# DBMATE_MIGRATIONS_DIR is a folder with migrations for this project it is migrations
DBMATE_MIGRATIONS_DIR=migrations
# DATABASE_URL is your database connection string
DATABASE_URL=postgres://postgres:postgres@db:5432/movie_games?sslmode=disable


# database
# Here are your database credentials
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=movie_games
```
4. Edit `nginx.conf` file to suit your needs
5. Run `docker-compose -f docker-compose.yml.local build` to build your containers
6. Run `docker-compose -f docker-compose.yml.local up -d` to run your project
7. It will automatically assign telegram bot callbacks, but you need to manually
   set your bot menu button with @BotFather telegram bot  if you want to open web app
   with menu button.

</details>

## Run everything yourself

For development/debugging you can run everything separately.

### Setup Database

<details><summary><b>Show instructions</b></summary>

1. Create your <a href=https://www.postgresql.org/>PostgreSQL</a> database
2. Install <a href=https://github.com/amacneil/dbmate#dbmate>Dbmate</a>
3. In the backend folder set your .env file with 
   <a href=https://github.com/amacneil/dbmate#dbmate>Dbmate</a> 
   Environment Variables
```dotenv
# DBMATE_NO_DUMP_SCHEMA is boolean for createing dump schema file you can leave it 
# to true, as dump schema is not needed for this project
DBMATE_NO_DUMP_SCHEMA=true
# DBMATE_MIGRATIONS_DIR is a folder with migrations for this project it is migrations
DBMATE_MIGRATIONS_DIR=migrations
# DATABASE_URL is your database connection string
DATABASE_URL=postgres://postgres:postgres@db:5432/movie_games?sslmode=disable
```
4. Run `dbmate up` to apply migrations
</details>

### Setup Backend

<details><summary><b>Show instructions</b></summary>

1. Setup <a href=https://go.dev/>Go</a> 1.20 or newer
2. Go to the backend folder
3. Export your Environment Variables
```dotenv
# REST_LISTEN is adress at which web server will listen to requests
REST_LISTEN=0.0.0.0:8888
# DATABASE is your database DSN string
DATABASE='user=postgres password=postgres host=localhost port=5432 dbname=movie_games sslmode=disable'
# JWT_TOKEN is your JWT token secret you can write here any combination of symbols
JWT_TOKEN=sfhjahkfg8749GHGJHgjhds
# TELEGRAM_BOT_TOKEN is your Telegram bot token you can get it from @BotFather 
# when registering your telegram bot
TELEGRAM_BOT_TOKEN=123456789:qwertyuioASDFGHJKLzxcvbnm
# FRONTEND_HOSTNAME is your frontend address
FRONTEND_HOSTNAME=localhost:8080
# BACKEND_HOSTNAME is your backend address
BACKEND_HOSTNAME=localhost:8888
```
4. Run `go mod download` to download all the dependencies 
5. Run `go build main.go` this will run your backend web server

</details>

### Setup Frontend

<details><summary><b>Show instructions</b></summary>

1. Install <a href=https://nodejs.org/en>Node.js</a>
2. Go to the frontend folder
3. Run `npm install` to install all the dependencies
4. Run `npm run serve` to serve frontend on localhost. You can also use
   `npm run dev` which will build dist and auto update dist with all the changes 
   and the serve it with a <a href=https://www.npmjs.com/package/serve>serve</a> package

</details>

### Open app in Telegram

You have 2 options here:
* Create a test environment telegram account and run app from localhost 
  in test environment using `npm run serve`
* Expose your local ports at https address using utility like
  <a href=https://ngrok.com/>ngrok</a> and test app through your normal telegram
  account

More info on the 2nd method:

<details><summary><b>Show instructions</b></summary>

1. Register an <a href=https://ngrok.com/>ngrok</a> account
2. Setup ngrok config to tunnel 2 different ports
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
3. Change your backend Environment Variables to addresses given to you by ngrok.
   For example:
```dotenv
FRONTEND_HOSTNAME=https://2acf-188-233-88-176.ngrok-free.app
BACKEND_HOSTNAME=https://7cc1-188-233-88-176.ngrok-free.app
```
4. Run backend with `go build main.go` in the backend folder. It will run at 8888 port by default,
   or you can change it with `REST_LISTEN` Environment Variable
5. Change your frontend Environment Variable to address given to you by ngrok.
   For example:
```dotenv
# VUE_APP_BASE_URL is your backend address, requests to backend will go there
VUE_APP_BASE_URL=https://7cc1-188-233-88-176.ngrok-free.app
```
6. Run frontend with `npm run dev` to generate dist
7. Serve frontend with <a href=https://www.npmjs.com/package/serve>serve</a>
   by running `serve -l 8080` in the frontend folder
8. Register frontend ngrok address as app menu button url at @BotFather 
9. Now you can develop/debug app. More on debugging telegram apps you can see here
   <a href=https://core.telegram.org/bots/webapps#testing-mini-apps>Testing Mini Apps</a>
</details>

***

Made in 2 weeks for Telegram Mini Apps Contest.