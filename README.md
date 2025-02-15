# REST API for planning your day
## Description
MyDay API is a RESTful API for day planning. You can use it to add, edit and delete tasks and events. The API supports standart HTTP methods and returns data in JSON format. 
## Technologies and frameworks
  - API in accordance with REST principles.
  - The structure of the application in accordance with the principles of the <b>Clean Architecture</b>.
  - Storing data using <b>Postgres</b>. Generation of migration files.
  - HTTP server <a href = https://github.com/go-chi/chi>go-chi/chi</a>.
  - Configuration using <a href = https://github.com/ilyakaznacheev/cleanenv>ilyakaznacheev/cleanenv</a>. Working with environment variables.
  - Implemented registration and authentication using <a href = https://github.com/gorilla/sessions>gorilla/sessions</a>.
  - Writing SQL queries using <a href = https://github.com/jmoiron/sqlx>sqlx</a>.
  - Working with Dockerfile and docker-compose.
  - API documentation using <a href = https://github.com/go-swagger/go-swagger>swagger</a>.
## Prerequisites
Before you begin, ensure you have the following installed on your machine:
  - Go(version 1.23 or higher)
  - Docker
## Installation 
### 1. Clone the repository
```
git clone https://github.com/NikitaBorovikov/my-day.git
cd my-day
```
### 2. Set up environment variables
Create a ```.env``` file in the root of the project:
```
cp .env.example .env
```
Open the .env file and fill in the required values:
```
PG_HOST=postgres
PG_PORT=5432
PG_USER=your_user
PG_PASSWORD=your_password
PG_NAME=your_db_name
SESSION_KEY=your_session_key
```
Replace ```your_user```, ```your_password```, ```your_db_name```, ```your_session_key``` with your actual values.

### 3. Run db migration
If the application is being launched for the first time, migrations must be applied to the database:
```
make migrate
```
### 4. Build and run the application:
Build the docker image:
```
make docker-image 
```
Run the application:
```
make run
```




  
