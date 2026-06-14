## Project overview

`tekd-be` is the backend API for Tekd, built with Go and the Gin web framework. It exposes versioned HTTP endpoints under `/api/v1`.

## Tech stack

- **Language**: Go 1.26
- **Framework**: Gin
- **Config**: Environment variables via `godotenv` (`.env` or system env)
- **Module**: `github.com/akshara-devs/tekd-be`

## Project structure

```
main.go                 # Entry point — loads config, starts HTTP server
pkg/
  config.go             # Env-based configuration
  api.go                # Shared HTTP response helpers
router/
  route.go              # Router setup and middleware
  api/v1/               # Versioned handlers (one file per route group)
```

## Conventions

- **Routing**: Register routes in `router/route.go` under the `/api/v1` group. Add handlers in `router/api/v1/`.
- **API responses**: Use `pkg.JSON()` for all JSON responses. It returns a consistent shape:

  ```go
  pkg.JSON(c, http.StatusOK, "message", data)
  // { "statusCode": 200, "message": "message", "data": { ... } }
  ```

- **Config**: Load settings through `pkg.LoadConfig()`. Do not hardcode ports, timeouts, or run mode.
- **Middleware**: Keep `gin.Logger()` and `gin.Recovery()` on the root router unless there is a specific reason to change them.
- **Packages**: Put shared, cross-cutting code in `pkg/`. Keep route handlers thin; move business logic into dedicated packages as the project grows.

## Common commands

```bash
go run .                # Start dev server (default port 8080)
go build -o tekd-be .   # Build binary
go test ./...           # Run all tests
go mod tidy             # Sync dependencies after imports change
```

## Environment

Create a `.env` file at the project root (gitignored). Supported variables:

| Variable        | Default | Description                    |
|-----------------|---------|--------------------------------|
| `HTTP_PORT`     | `8080`  | Server listen port             |
| `RUN_MODE`      | `debug` | Gin mode (`debug` / `release`) |
| `READ_TIMEOUT`  | `10`    | Read timeout in seconds        |
| `WRITE_TIMEOUT` | `10`    | Write timeout in seconds       |

## Git workflow

Always use [Conventional Commits](https://www.conventionalcommits.org/) for commit messages.

**Format**: `<type>[optional scope]: <description>`

| Type       | Use for                                      |
|------------|----------------------------------------------|
| `feat`     | New feature or endpoint                      |
| `fix`      | Bug fix                                      |
| `refactor` | Code change that neither fixes nor adds      |
| `chore`    | Tooling, deps, config (no production change) |
| `docs`     | Documentation only                           |
| `test`     | Adding or updating tests                     |

**Examples**:

```
feat(api): add user registration endpoint
fix(config): parse HTTP_PORT when env value is invalid
chore(deps): bump gin to v1.12.0
docs: add AGENTS.md with project conventions
```

- Use imperative mood in the subject line (`add`, not `added`).
- Keep the subject under 72 characters.
- Add a body after a blank line when the change needs context.
- Only create commits when explicitly asked.
- Never commit `.env` or other secrets.

## When adding features

1. Add handler(s) in `router/api/v1/`.
2. Register the route in `router/route.go`.
3. Return responses via `pkg.JSON()`.
4. Add shared logic to `pkg/` or a new package if it grows beyond a single handler.
5. Run `go build .` and `go test ./...` before finishing.

## Scope

- Keep changes focused; match existing patterns in neighboring files.
- Only add tests when requested or when they provide meaningful coverage.
- Do not edit generated files or commit secrets.
