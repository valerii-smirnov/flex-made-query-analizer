# postgres query analyzer application

---

To run application and database with preinstalled pg_stat_statements and preloaded test database 
dump use the command `make run` from the root project folder. Application api will be 
available on `http://127.0.0.1:8080`. Swagger file you will find in the root of the 
project

---

All necessary dependencies, data and dumps are configured automatically. You don't need to do anything manually.

---

Database is available by address `127.0.0.1:5432`. Work database name is `dvdrental`. 
Login and password for database you'll find in `docker/.env` file.

---

After database dump restore some queries will be available, but if you want to execute your custom queries,
please connect to the database and execute your queries. After execution these queries
will be available in application api response.

---

To stop application and database containers you can use `make stop` command from the
root of the project.

---

To run unit test with coverage run command `make test` from the root of the project.

---

Every command and application is dockerized, and you don't need to have installed soft
on your local machine. Only requirement is you need to have installed `docker` and 
`docker-compose`.