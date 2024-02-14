
# currencyAPI

  

## Running Locally

1. Clone the repository:

  

```bash
git  clone  https://github.com/viciousvs/currencyAPI.git

cd  ccurrencyAPI/
```

2. Edit .env file

3. Run Postgresql from docker

```bash
docker  run  --name  currencyDB  -p  5432:5432  -e  POSTGRES_USER=admin  -e  POSTGRES_PASSWORD=admin  -e  POSTGRES_DB=currency  -d  postgres
```

4. Download dependencies

```bash
go  mod  download
```

5. Run server

```bash
go  run  cmd/app/main.go
```

6. Run cron
```bash
go run cmd/cron/main.go
```

## Endpoints

### Get Currencies: [http://localhost:8080/currencies](http://localhost:8080/currencies)
``` bash
curl http://localhost:8080/currencies
```
### Get Currency by rate name: [http://localhost:8080/currencies/KZT](http://localhost:8080/currencies/KZT)
``` bash
curl http://localhost:8080/currencies/KZT
```
### Update: [http://localhost:8080/update](http://localhost:8080/update)
``` bash
curl http://localhost:8080/update
```

## TODO
- write unit tests
- makefile
- swagger docs
- docker-compose
- ci/cd
- frontend client