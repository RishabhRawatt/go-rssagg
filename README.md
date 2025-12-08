# go-rssagg

A Go-based RSS feed aggregator that fetches, stores, and serves RSS feeds through a RESTful API.

## Features

- 📰 RSS/Atom feed parsing and aggregation
- 🔄 Automatic feed scraping with configurable intervals
- 👤 User management with API key authentication
- 📝 Post management and retrieval
- 🔐 API key-based authentication
- 🗄️ PostgreSQL database storage
- 🌐 RESTful API with JSON responses

## Prerequisites

- Go 1.21 or higher
- PostgreSQL database
- sqlc (for generating database code)
- goose (for database migrations)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/RishabhRawatt/go-rssagg.git
cd go-rssagg
```

2. Install dependencies:
```bash
go mod download
```

3. Install development tools:
```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install github.com/pressly/goose/v3/cmd/goose@latest
```

4. Set up your environment variables:
```bash
cp .env.example .env
# Edit .env with your database credentials
```

5. Run database migrations:
```bash
make migrate-up
```

## Configuration

Create a `.env` file in the root directory with the following variables:

```env
PORT=8080
DB_URL=postgres://username:password@localhost:5432/rssagg?sslmode=disable
```

## Running the Application

### Development
```bash
make run
```

### Production
```bash
make build
./go-rssagg
```

## API Endpoints

### Health Check
- `GET /v1/healthz` - Check if the server is running
- `GET /v1/err` - Test error handling

### Users
- `POST /v1/users` - Create a new user
  ```json
  {
    "name": "John Doe"
  }
  ```
  Returns user object with API key

- `GET /v1/users` - Get current user (requires authentication)
  - Header: `Authorization: ApiKey {your_api_key}`

### Feeds
- `POST /v1/feeds` - Create a new feed (requires authentication)
  ```json
  {
    "name": "Blog Name",
    "url": "https://example.com/feed.xml"
  }
  ```

- `GET /v1/feeds` - Get all feeds

### Feed Follows
- `POST /v1/feed_follows` - Follow a feed (requires authentication)
  ```json
  {
    "feed_id": "uuid-here"
  }
  ```

- `GET /v1/feed_follows` - Get all feed follows for current user (requires authentication)

- `DELETE /v1/feed_follows/{feedFollowID}` - Unfollow a feed (requires authentication)

### Posts
- `GET /v1/posts` - Get posts from followed feeds (requires authentication)
  - Returns the 10 most recent posts from feeds you follow

## Authentication

The API uses API key authentication. Include your API key in the request header:

```
Authorization: ApiKey {your_api_key}
```

You receive an API key when creating a user account.

## Database Schema

The application uses PostgreSQL with the following tables:

- `users` - User accounts with API keys
- `feeds` - RSS feed sources
- `feed_follows` - User feed subscriptions
- `posts` - Aggregated posts from feeds

## Background Tasks

The application runs a background scraper that:
- Fetches feeds every minute (configurable)
- Runs 10 concurrent scrapers
- Updates the database with new posts
- Tracks the last fetch time for each feed

## Development

### Generate Database Code
After modifying SQL queries or schema:
```bash
make sqlc
```

### Database Migrations
Create new migrations in `sql/schema/` following the naming pattern:
```
NNN_description.sql
```

Apply migrations:
```bash
make migrate-up
```

Rollback migrations:
```bash
make migrate-down
```

## Project Structure

```
go-rssagg/
├── main.go                 # Application entry point
├── handler_*.go            # HTTP request handlers
├── middleware_auth.go      # Authentication middleware
├── models.go              # Data models
├── scraper.go             # Feed scraping logic
├── rss.go                 # RSS parsing
├── json.go                # JSON response helpers
├── internal/
│   ├── auth/              # Authentication utilities
│   └── database/          # Generated database code (sqlc)
├── sql/
│   ├── schema/            # Database migrations
│   └── queries/           # SQL queries for sqlc
├── .env.example           # Example environment configuration
├── sqlc.yaml              # sqlc configuration
└── Makefile              # Development commands
```

## Technologies Used

- **Go** - Programming language
- **Chi** - HTTP router
- **PostgreSQL** - Database
- **sqlc** - SQL compiler (generates type-safe Go code)
- **goose** - Database migration tool
- **godotenv** - Environment variable management

## License

MIT License

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
