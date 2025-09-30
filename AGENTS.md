# Repository Guidelines

## Project Structure & Module Organization
The repository is a collection of focused Go modules grouped by purpose. Use `experiment/` for architectural and performance spikes, `library_selection/` for side-by-side evaluations of frameworks and ORMs, `knowledge/` for language idioms and patterns, and `how2use/` for minimal, ready-to-run feature demos. The `db/` directory provides shared Docker Compose assets, while `try/` and `olds/` keep legacy prototypes that inform current work—treat them as reference only. Each leaf module owns its own `go.mod`, so keep new code scoped to the directory where it belongs.

## Build, Test, and Development Commands
Run Go commands from the module root you are touching: `go build ./...` validates compilation, `go test ./...` executes unit tests, and `go run .` drives runnable samples. When a directory exposes a Makefile (for example `library_selection/db_migration/sqldef/Makefile`), prefer those targets—`make mysql-start` boots the local MySQL container and `make go-run` executes the sample with matching env vars. Use `docker compose -f db/docker-compose.yml up -d` to share infrastructure across experiments.

## Coding Style & Naming Conventions
Follow idiomatic Go style: tab-indented blocks, short lower_snake module paths, and mixedCaps identifiers for exported symbols. Always run `gofmt` and `goimports` before committing; lint with `golangci-lint run ./...` when available. Keep package names singular and concise, and align filenames with the primary type (e.g., `user_service.go` for `UserService`).

## Testing Guidelines
Unit tests live alongside source files with `_test.go` suffixes and table-driven structure. Name tests using `TestScenario_Action` to clarify intent, and favor subtests for variants. For integration cases (database, network), gate execution behind build tags or environment checks and document prerequisites in a README. Use `go test -cover ./...` to confirm coverage before raising a PR; modules that touch external services should note required containers in their README.

## Commit & Pull Request Guidelines
Commit messages begin with an imperative summary (<72 chars) and optionally include a descriptive body that references impacted paths (see `git log`). Group related changes per commit and avoid mixing experiments. Pull requests need a succinct problem statement, testing notes, and links to relevant issues. Include screenshots or logs when behavior changes, and call out any new Compose services or environment variables.
