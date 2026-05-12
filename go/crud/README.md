# Simple CRUD Application in Go

This is a simple CRUD (Create, Read, Update, Delete) web application built with Go's standard library `net/http` package.

## Features

- Create new items
- Read all items or a specific item
- Update existing items
- Delete items
- In-memory storage (data is lost when server restarts)
- JSON API endpoints

## API Endpoints

| Method | Endpoint        | Description               |
|--------|-----------------|---------------------------|
| GET    | `/items`        | Get all items             |
| GET    | `/items/{id}`   | Get a specific item       |
| POST   | `/items`        | Create a new item         |
| PUT    | `/items/{id}`   | Update an existing item   |
| DELETE | `/items/{id}`   | Delete an item            |

## Item Structure

Each item has the following fields:
- `ID`: Unique identifier (string)
- `Name`: Name of the item (string)
- `Description`: Description of the item (string)

## How to Run

1. Make sure you have Go installed (version 1.16 or later)
2. Navigate to the `crud` directory:
   ```bash
   cd go/crud
   ```
3. Run the application:
   ```bash
   go run main.go
   ```
4. The server will start on `http://localhost:8080`

## Testing the API

You can test the API using curl or tools like Postman:

### Get all items
```bash
curl http://localhost:8080/items
```

### Get a specific item
```bash
curl http://localhost:8080/items/1
```

### Create a new item
```bash
curl -X POST http://localhost:8080/items \
  -H "Content-Type: application/json" \
  -d '{"name":"Laptop","Description":"Gaming laptop"}'
```

### Update an item
```bash
curl -X PUT http://localhost:8080/items/1 \
  -H "Content-Type: application/json" \
  -d '{"id":"1","name":"Laptop Baru","Description":"Laptop gaming terbaru"}'
```

### Delete an item
```bash
curl -X DELETE http://localhost:8080/items/1
```

## Project Structure

- `main.go`: Contains the complete CRUD application with:
  - Item struct definition
  - In-memory store with mutex for concurrency safety
  - HTTP handlers for all CRUD operations
  - Main function to start the server

## Notes

- This application uses in-memory storage, so data will be lost when the server restarts
- For production use, consider replacing the in-memory store with a database
- The ID generation is simple (using Unix nanoseconds) and not guaranteed to be unique in high-concurrency scenarios
- Error handling is basic but demonstrates the concepts