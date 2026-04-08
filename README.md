# Go Generic API

A simple REST API built with Go and the Gin framework.
This project serves as a personal exercise to practice building web services, routing, models, and handlers in Go.

## Endpoints

| Method | Rute              | Description                     |
|--------|-------------------|---------------------------------|
| GET    | /api/items        | Obtiene todos los items         |
| GET    | /api/items/:id    | Obtiene un item por su ID       |
| POST   | /api/items        | Crea un nuevo item              |
| DELETE | /api/items/:id    | Elimina un item por su ID       |

## Run the Project

From the project root:
go run cmd/server/main.go

## Project structure

go-api/
 ├── cmd/
 │    └── server/
 │         └── main.go
 ├── internal/
 │    ├── routes/
 │    │     └── router.go
 │    ├── handlers/
 │    │     └── items.go
 │    └── models/
 │          └── item.go
 ├── go.mod
 ├── go.sum
 └── README.md


## How it works

### cmd/server/main.go

Entry point of the application.

Loads a fully configured router from internal/routes.

Starts the HTTP server on port 8080.

### internal/routes/router.go

Defines the routing configuration.

gin.Default() creates a router with logging and recovery middleware.

Groups all endpoints under /api.

Maps each route to its corresponding handler.

### internal/models/item.go

Defines the Item data structure.

### internal/handlers/items.go

Contains the business logic of the API.

Maintains an in‑memory list of items (simulating a database).

GetItems → returns all items.

GetItem → returns a specific item by ID or 404 if not found.

CreateItem → parses JSON, assigns an ID, and adds the item to the list.

DeleteItem → removes an item by ID or returns 404 if not found.

## Notes

The API uses an in-memory data store. Restarting the server resets all data