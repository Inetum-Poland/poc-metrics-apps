# PoC - Metrics, Tracing, Logging, and Observability

## Concept

The main goal of this PoC is to demonstrate how to use OpenTelemetry to collect metrics, tracing, and logging data from different programming languages and frameworks.

### Languages

- .NET
- Go
- Python

## API Endpoints

### `GET /api/long_run`

### `GET /api/short_run`

### `GET /api/database_run`

## Main Functions

### `longRun()`

### `shortRun()`

### `databaseRun()`

## Functions

### `add(int a, int b)`

### `subtract(int a, int b)`

### `multiply(int a, int b)`

### `divide(int a, int b)`

### `readData(string query)`

## Database Schema

| Collection | Field | Type     | Description                        |
| ---------- | ----- | -------- | ---------------------------------- |
| Data       | _id   | ObjectId | Unique identifier for the document |
| Data       | data  | int      | Data field                         |

## Docker Compose and Makefile

To run the application, you can use the provided `docker-compose` file or `makefile` which include some automation.

```bash
> make docker-compose
```

## Prerequisites

- Docker
- Docker Compose
- Make

## Links

- [OpenTelemetry](https://opentelemetry.io/)
- [OpenTelemetry .NET](https://opentelemetry.io/docs/languages/net/instrumentation/)
- [OpenTelemetry Go](https://opentelemetry.io/docs/languages/go/instrumentation/)
- [OpenTelemetry Python](https://opentelemetry.io/docs/languages/python/instrumentation/)
