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

### FizzBuzz Exercises (Introduction)

| Exercise         | File                              | Description                             |
| ---------------- | --------------------------------- | --------------------------------------- |
| FizzBuzz Basic   | `exercise/fizzbuzz/fizzbuzz.go`   | Implement the classic FizzBuzz function |
| FizzBuzz Multi   | `exercise/fizzbuzz/multi.go`      | Process a slice of numbers              |
| FizzBuzz Map     | `exercise/fizzbuzz/multimap.go`   | Return results as a map                 |
| FizzBuzz Struct  | `exercise/fizzbuzz/fizzstruct.go` | Use structs and methods                 |
| FizzBuzz Handler | `exercise/fizzbuzz/server.go`     | HTTP handler for FizzBuzz               |

### API/NetHTTP Exercises

| Exercise        | File                               | Description                |
| --------------- | ---------------------------------- | -------------------------- |
| Database Setup  | `exercise/api/nethttp/database.go` | Set up database connection |
| Request Reading | `exercise/api/nethttp/request.go`  | Read from HTTP request     |
| Result Writing  | `exercise/api/nethttp/result.go`   | Write JSON responses       |
| List Handler    | `exercise/api/nethttp/list.go`     | List all pets endpoint     |
| Get By ID       | `exercise/api/nethttp/getById.go`  | Get single pet endpoint    |
| Update Handler  | `exercise/api/nethttp/update.go`   | Update pet endpoint        |

### Database Exercises

| Exercise    | File                               | Description                      |
| ----------- | ---------------------------------- | -------------------------------- |
| SQL Select  | `exercise/database/sql/select.go`  | SELECT queries with database/sql |
| SQLx Select | `exercise/database/sqlx/select.go` | SELECT queries with sqlx         |

## Running Tests

### Test your exercise implementation

```bash
# Test a specific exercise
go test ./exercise/fizzbuzz/... -run ^TestShouldReturnSameNumber$

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
