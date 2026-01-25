# Go Introduction + API - Exercises

This repository contains the exercises for Module 1 of the Go Training course.

## Prerequisites

- Go 1.25.6 or later
- PostgreSQL (for database exercises)
- Docker (optional, for PostgreSQL setup)

## Repository Structure

```
go-introduction-api/
├── exercise/           # Exercise stubs - implement these!
├── solution/           # Complete solutions
├── examples/           # Demo code and working examples
│   ├── api/cmd/petshelter/  # Complete API with SwaggerUI
│   ├── database/       # Database examples (bob, gorm, sqlx, connect)
│   └── ...
├── shelter/            # Shared types
└── docker/             # Database setup
```

## Exercise Order

### Pet Exercises (Introduction)

| Exercise          | File                          | Description                             |
| ----------------- | ----------------------------- | --------------------------------------- |
| Categorize Age    | `exercise/pets/categorize.go` | Categorize pet age (Puppy/Adult/Senior) |
| Filter Adult Ages | `exercise/pets/slice.go`      | Filter ages in the Adult range          |
| Count by Category | `exercise/pets/map.go`        | Count occurrences using maps            |
| Pet Info Struct   | `exercise/pets/petstruct.go`  | Use structs and methods                 |

### Database Exercises

| Exercise   | File                              | Description                      |
| ---------- | --------------------------------- | -------------------------------- |
| SQL Select | `exercise/database/sql/select.go` | SELECT queries with database/sql |

### API/NetHTTP Exercises

| Exercise        | File                              | Description             |
| --------------- | --------------------------------- | ----------------------- |
| Result Writing  | `exercise/api/nethttp/result.go`  | Write JSON responses    |
| Request Reading | `exercise/api/nethttp/request.go` | Read from HTTP request  |
| List Handler    | `exercise/api/nethttp/list.go`    | List all pets endpoint  |
| Get By ID       | `exercise/api/nethttp/getById.go` | Get single pet endpoint |
| Update Handler  | `exercise/api/nethttp/update.go`  | Update pet endpoint     |



## Running Tests

### Test your exercise implementation

```bash
# Test a specific exercise
go test ./exercise/pets/... -run ^TestCategorizeAge$

# Test all exercises (will fail until implemented)
go test ./exercise/...
```

### Verify against solutions

```bash
# All solution tests should pass
go test ./solution/...
```

## Database Setup

For database exercises, you need a PostgreSQL instance. Use the provided Docker setup:

```bash
cd docker
docker-compose up -d
```

See `docker/README.md` for more details.

## Go Version

This module requires Go 1.25.6 or later. Check your version:

```bash
go version
```
