# Ticket System API

A REST API for a ticket management system built with Go, PostgreSQL, and JWT authentication.

## Status Flow

```
open -> in_progress -> closed
```

A closed ticket cannot be reopened or moved back to `in_progress`.

## Endpoints

| Method | Endpoint                | Description                    | Auth |
|--------|-------------------------|--------------------------------|------|
| GET    | /health                 | Health check                   | No   |
| POST   | /auth/register          | Register a new user            | No   |
| POST   | /auth/login             | Login and receive JWT token    | No   |
| POST   | /tickets                | Create a new ticket            | Yes  |
| GET    | /tickets                | List all tickets for user      | Yes  |
| GET    | /tickets/{id}           | Get a specific ticket by ID    | Yes  |
| PATCH  | /tickets/{id}/status    | Update ticket status           | Yes  |

## Local Run

```bash
with docker-compose
docker-compose up --build

# Or directly
go run ./cmd/api
```

## Docker

```bash
docker build -t ticket-system .
docker run -p 8080:8080 ticket-system
```

## Deployment

Deployed URL: https://ticket-system-xi-nine.vercel.app

## Environment Variables

Copy `.env.example` to `.env` and update values.

| Variable    | Default       | Description          |
|-------------|---------------|----------------------|
| JWT_SECRET  | default-dev-secret-change-in-prod | JWT signing secret |
| DB_HOST     | localhost     | PostgreSQL host      |
| DB_PORT     | 5432          | PostgreSQL port      |
| DB_USER     | postgres      | PostgreSQL user      |
| DB_PASSWORD | postgres      | PostgreSQL password  |
| DB_NAME     | ticketdb      | PostgreSQL database  |
| SERVER_ADDR | :8080         | Server address       |

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

# Create ticket
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

- JWT tokens expire after 24 hours.
- Passwords are stored as bcrypt hashes.
- Users can only view and update their own tickets.
- No admin role or ticket assignment is implemented.
- Default JWT secret is not suitable for production.
