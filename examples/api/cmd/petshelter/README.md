# Pet Shelter API

REST API server for managing pets in a shelter.

## Prerequisites

Start the PostgreSQL database:

```bash
cd docker && docker-compose up -d
```

## Run

```bash
go run ./examples/api/cmd/petshelter
```

With custom port:

```bash
PORT=3000 go run ./examples/api/cmd/petshelter
```

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | /pet | List all pets |
| GET | /pet/{id} | Get a pet by ID |
| POST | /pet/{id}/rename | Rename a pet |
| GET | /swagger | SwaggerUI |
| GET | /openapi.yaml | OpenAPI specification |
