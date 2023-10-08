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

# Usage

## Docker compose

You can run this app with
<a href=https://github.com/Sergey-pr/movie-games-tg/blob/main/docker-compose.yml>
docker-compose.yml</a> file.

<details><summary><b>Show instructions</b></summary>

1. Make a copy of
   <a href=https://github.com/Sergey-pr/movie-games-tg/blob/main/docker-compose.yml>
   docker-compose.yml</a> and name it docker-compose.yml.local
2. Edit **Environment Variables** to suit your needs

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

3. Run `docker-compose -f docker-compose.yml.local build` to build your containers
4. Run `docker-compose -f docker-compose.yml.local up -d` to run your project
5. It will automatically assign telegram bot callbacks, but you need to manually
   set your bot menu button with @BotFather telegram bot  if you want to open web app
   with menu button.
</details>

## Run everything yourself

For development/debugging you can run everything separately.

### Setup Database

<details><summary><b>Show instructions</b></summary>

1. Create your postgres database
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
2. Export your Environment Variables
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
3. Run `go mod download` to download all the dependecies
4. Run `go build main.go` this will run your backend web server

</details>

### Setup Frontend

<details><summary><b>Show instructions</b></summary>

1. Install <a href=https://nodejs.org/en>Node</a>
2. Go to frontend folder
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
8. Register frontend address as app url at @BotFather 
</details>
