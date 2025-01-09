# Observer-Logs-Backend

To start the server, run the following command inside the directory:

```bash
go run .
```

The server will default to listening on port `8080`, with the following routes available:

## Available Routes

- **GET /query**: Fetches data, which can be filtered using query parameters:
  - `start`: Start time (Unix timestamp)
  - `end`: End time (Unix timestamp)
  - `text`: Text to search for in the log body

- **POST /ingest**: Adds data to memory.

- **GET /refresh**: Refreshes the data in memory.
