# PoC - Metrics, Tracing, Logging, and Observability

## Concept

The main goal of this PoC is to demonstrate how to use OpenTelemetry to collect metrics, tracing, and logging data from different programming languages and frameworks.

### Languages

- GoLang
- .NET
- Python

## API Endpoints

### `GET /api/long_run`

- [golang](http://localhost:8080/api/long_run)
- [dotnet](http://localhost:8081/api/long_run)
- [python](http://localhost:8082/api/long_run)

This should call the `longRun()` function.

### `GET /api/short_run`

- [golang](http://localhost:8080/api/short_run)
- [dotnet](http://localhost:8081/api/short_run)
- [python](http://localhost:8082/api/short_run)

This should call the `shortRun()` function.

### `GET /api/database_run`

- [golang](http://localhost:8080/api/database_run)
- [dotnet](http://localhost:8081/api/database_run)
- [python](http://localhost:8082/api/database_run)

This should call the `databaseRun()` function.

### `GET /api/failed_run`

- [golang](http://localhost:8080/api/failed_run)
- [dotnet](http://localhost:8081/api/failed_run)
- [python](http://localhost:8082/api/failed_run)

This should call the `failedRun()` function.

## Main Functions

### `longRun()`

This function should call the following:

- `add(int a, int b)`
- `subtract(int a, int b)`
- `multiply(int a, int b)`
- `divide(int a, int b)`

### `shortRun()`

This function should sleep for 100 milliseconds.

### `databaseRun()`

This function should connect to the database, find all documents, and return the result.

### `failedRun()`

This function should sleep for 100 milliseconds and then fail.

## Database Schema

| Collection | Field        | Type     | Description                                   |
| ---------- | ------------ | -------- | --------------------------------------------- |
| Data       | `_id`        | ObjectId | (internal) Unique identifier for the document |
| Data       | `created_at` | Time     | (internal) N/a                                |
| Data       | `updated_at` | Time     | (internal)N/a                                 |
| Data       | `data`       | int      | Data field                                    |

## Docker Compose and Makefile

To run the application, you can use the provided `docker-compose` file or `makefile` which include some automation.

```bash
> make docker-compose
```

### Docker Compose Links

| Service       | Links                  |
| ------------- | ---------------------- |
| Grafana       | http://localhost:3000  |
| Prometheus    | http://localhost:9090  |
| Loki          | http://localhost:3100  |
| Tempo         | http://localhost:3200  |
| Mongo-Express | http://localhost:27018 |

## Prerequisites

- Docker
- Docker Compose
- Make

## Links

- [OpenTelemetry](https://opentelemetry.io/)
- [OpenTelemetry .NET](https://opentelemetry.io/docs/languages/net/instrumentation/)
- [OpenTelemetry Go](https://opentelemetry.io/docs/languages/go/instrumentation/)
- [OpenTelemetry Python](https://opentelemetry.io/docs/languages/python/instrumentation/)
