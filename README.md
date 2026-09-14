# Ticket System API

docker build -t ticket-system . && docker run -p 8080:8080 ticket-system
OR 
go run ./cmd/api

A simple REST API for a ticket management system built with Go, SQLite, and JWT authentication.

## Features

- User registration and login with JWT authentication
- Create, list, and view tickets
- Ownership-based authorization (users can only see their own tickets)
- Ticket status workflow: `open` -> `in_progress` -> `closed`
- SQLite persistence

## API Endpoints

| Method | Endpoint                | Description                    | Auth |
|--------|-------------------------|--------------------------------|------|
| GET    | /health                 | Health check                   | No   |
| POST   | /auth/register          | Register a new user            | No   |
| POST   | /auth/login             | Login and receive JWT token    | No   |
| POST   | /tickets                | Create a new ticket            | Yes  |
| GET    | /tickets                | List all tickets for user      | Yes  |
| GET    | /tickets/{id}           | Get a specific ticket by ID    | Yes  |
| PATCH  | /tickets/{id}/status    | Update ticket status           | Yes  |

## Status Flow

```
open -> in_progress -> closed
```

A closed ticket cannot be reopened or moved back to in_progress.

## Local Run

```bash
# Without Docker
go run ./cmd/api

# With Docker
docker build -t ticket-system .
docker run -p 8080:8080 ticket-system
```

## Environment Variables

| Variable    | Default                | Description          |
|-------------|------------------------|----------------------|
| JWT_SECRET  | default-dev-secret...  | Secret for JWT signing |
| DB_PATH     | ./tickets.db           | Path to SQLite file   |

## Example Usage

```bash
# Register
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"secret123"}'

# Login
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"secret123"}'

# Create ticket (use token from login response)
curl -X POST http://localhost:8080/tickets \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"title":"My ticket","description":"Some description"}'

# List tickets
curl http://localhost:8080/tickets \
  -H "Authorization: Bearer <token>"

# Update status
curl -X PATCH http://localhost:8080/tickets/<id>/status \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"status":"in_progress"}'
```

## Assumptions

- SQLite is used as the persistent store (file-based, no external DB needed).
- JWT tokens expire after 24 hours.
- Default JWT secret is used if `JWT_SECRET` is not set (not suitable for production).
- No admin role or ticket assignment is implemented per requirements.
