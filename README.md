# enlabs-assignment

run `docker-compose up --build`

application will be at 8080 port

application creates 3  test users by default

## routes

- POST /user
```
request:    { name: string }
response:   { user: { id: int, name: string, balance: float } }
```

- GET /user/:id

```
request:    {  }
response:   { user: { id: int, name: string, balance: float } }
```

- POST /user/:id/transaction

```
request:    { "state": "win"|"lose"|string, "amount": float, "transactionId": string }
response:   { user: { id: int, name: string, balance: float } }
```

## config

Set environment variables to configure service. Edit ./.env file for configuring service in docker compose
- ENLABS_DB_DSN - dsn for connecting to postgres (default: "host=localhost port=5432 user=postgres dbname=enlabs password=postgresPass")
- ENLABS_LOG_LEVEL - 0-6 - level of logs (default:0)
- ENLABS_LOG_FILE - file for logs (default: "" - use os.Stdout)
- ENLABS_SCHEDULERS_ENABLE - enable periodic tasks execution (default: false)
- ENLABS_SCHEDULERS_CANCELODDCONFIG_PERIOD - period of odd transactions cancellation task in minutes (default: 1)
- ENLABS_SCHEDULERS_CANCELODDCONFIG_NUMBERTOCANCEL - quanitity of tasks to cancel (defaut: 10)
