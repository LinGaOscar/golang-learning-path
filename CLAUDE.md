# CLAUDE.md

This is a personal Go learning-exercises repo: a flat collection of small, standalone `.go` programs used to practice core language concepts (variables, control flow, interfaces, error handling, concurrency). It is not a single cohesive application.

## Commands

Every file declares `package main` and its own `func main()`, so the module does **not** build as a whole — `go build ./...` / `go run .` will fail with "main redeclared in this block". Run or build exercises one file at a time:

```bash
go run main.go
go run variables.go
go run fizzbuzz.go
go run slice_growth.go
go run shapes.go
go run concurrency.go
go run db_errors.go
```

To compile a single exercise to a binary:

```bash
go build variables.go
./variables
```

No test files exist in this repo (no `_test.go`), so there is currently no `go test` target.

Module housekeeping:

```bash
go mod tidy
go fmt ./...
```

## Architecture

There is no architecture beyond a flat pile of independent example scripts at the repo root, each covering one Go topic:

- `main.go` — minimal hello-world entry point / environment check
- `variables.go` — variable declaration, constants, `iota`
- `fizzbuzz.go` — control flow (`for`, `switch`)
- `slice_growth.go` — slice length/capacity growth behavior
- `shapes.go` — structs, methods, interfaces (`Shape`, `Rectangle`, ...)
- `concurrency.go` — goroutines, `sync.WaitGroup`, `sync.Mutex`
- `db_errors.go` — custom error types, `errors.Is`/`errors.As`-style patterns

`QuickStart.md` is a Chinese-language study/review guide summarizing the concepts covered by the exercises (not code documentation). `go.mod` declares module `hello-golang`, Go 1.26.2.
