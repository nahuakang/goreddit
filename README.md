# GO Reddit Clone
Example based on Go Web Examples course.

## How to run locally
In three terminal sessions, run each of the following lines:
```sh
$ make postgres
$ make adminer
$ reflex -s go run cmd/goreddit/main.go
```

Login to adminer at `localhost:8080` and type in the following details:
```sh
System:   PostgreSQL
Server:   localhost
Username: postgres
Password: secret
Database: postgres
```

After logging into Adminer, run in the terminal the following command:
```sh
$ make migrate
```
