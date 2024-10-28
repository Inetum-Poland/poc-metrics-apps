# PoC - Metrics, Tracing, Logging, and Observability

## Concept

The main goal of this PoC is to demonstrate how to use OpenTelemetry to collect metrics, tracing, and logging data from different programming languages and frameworks.

### Languages

- .NET
- Go
- Python

## API Endpoints

### `GET /api/long_run`

This should call the `longRun()` function.

### `GET /api/short_run`

This should call the `shortRun()` function.

### `GET /api/database_run`

This should call the `databaseRun()` function.

### `GET /api/failed_run`

This should call the `failedRun()` function.

## Main Functions

### `longRun()`

This function should call the following:

- Add
- Substract
- Multiply
- Divide

### `shortRun()`

This function should sleep for 100 milliseconds.

### `databaseRun()`

This function should connect to the database, find all documents, and return the result.

### `failedRun()`

This function should sleep for 100 milliseconds and then fail.

## Helper Functions

### `add(int a, int b)`

### `subtract(int a, int b)`

### `multiply(int a, int b)`

### `divide(int a, int b)`

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
