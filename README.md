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
├── examples/           # Demo code (not exercises)
├── shelter/            # Shared types
├── docker/             # Database setup
└── openapi/            # API specification
```

## Exercise Order

### FizzBuzz Exercises (Introduction)

| # | Exercise | File | Description |
|---|----------|------|-------------|
| 1 | FizzBuzz Basic | `exercise/fizzbuzz/fizzbuzz.go` | Implement the classic FizzBuzz function |
| 2 | FizzBuzz Multi | `exercise/fizzbuzz/multi.go` | Process a slice of numbers |
| 3 | FizzBuzz Map | `exercise/fizzbuzz/multimap.go` | Return results as a map |
| 4 | FizzBuzz Struct | `exercise/fizzbuzz/fizzstruct.go` | Use structs and methods |
| 5 | FizzBuzz Handler | `exercise/fizzbuzz/server.go` | HTTP handler for FizzBuzz |

### API/NetHTTP Exercises

| # | Exercise | File | Description |
|---|----------|------|-------------|
| 6 | Database Setup | `exercise/api/nethttp/database.go` | Set up database connection |
| 7 | List Handler | `exercise/api/nethttp/list.go` | List all pets endpoint |
| 8 | Get By ID | `exercise/api/nethttp/getById.go` | Get single pet endpoint |
| 9 | Request Reading | `exercise/api/nethttp/request.go` | Read from HTTP request |
| 10 | Result Writing | `exercise/api/nethttp/result.go` | Write JSON responses |
| 11 | Update Handler | `exercise/api/nethttp/update.go` | Update pet endpoint |

### Database Exercises

| # | Exercise | File | Description |
|---|----------|------|-------------|
| 12 | Connect | `exercise/database/connect/connect.go` | Connect to PostgreSQL |
| 13 | SQL Select | `exercise/database/sql/select.go` | SELECT queries with database/sql |
| 14 | SQL Update | `exercise/database/sql/update.go` | UPDATE queries with database/sql |
| 15 | SQLx Select | `exercise/database/sqlx/select.go` | SELECT queries with sqlx |
| 16 | SQLBoiler | `exercise/database/sqlboiler/select.go` | ORM queries with SQLBoiler |

## Running Tests

### Test your exercise implementation

```bash
# Test a specific exercise
go test ./exercise/fizzbuzz/... -run TestShouldReturnSameNumber

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
