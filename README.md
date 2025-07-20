# Gator - RSS Feed Aggregator CLI

## Overview

Gator is a command-line RSS feed aggregator built with Go and PostgreSQL. It allows users to:
- Follow RSS feeds
- View aggregated content
- Manage their subscriptions

## Prerequisites

Before using Gator, ensure you have:

1. **Go** (1.20+) - [Installation Guide](https://go.dev/doc/install)
2. **PostgreSQL** (12+) - [Installation Guide](https://www.postgresql.org/download/)

## Installation

### 1. Install the Gator CLI

```bash
go install github.com/adibbelel/gator@latest
```

This will install the `gator` binary to your `$GOPATH/bin`.

### 2. Verify Installation

```bash
gator --version
```

## Configuration

### 1. Create Config File

Gator uses a JSON config file located at `~/.gatorconfig.json`:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": "your_username"
}
```

### 2. Database Setup

1. Create a PostgreSQL database:
   ```bash
   createdb gator
   ```

2. Run migrations:
   ```bash
   gator migrate up
   ```

## Usage

### Basic Commands

```bash
# Create Username
gator register <username>

# Login to Existing User
gator login <username>

# Add Feed
gator addfeed <feed_name> <feed_url>

# Follow a new feed
gator follow https://example.com/feed.xml

# List your followed feeds
gator feeds

# List articles from your feeds
gator articles

# Manage users
gator users
```

### Command Reference

| Command | Description |
|---------|-------------|
| `gator follow <url>` | Subscribe to an RSS feed |
| `gator unfollow <feed-id>` | Unsubscribe from a feed |
| `gator feeds` | List your subscriptions |
| `gator browse` | Browse through articles |
| `gator users` | Manage users |
| `gator migrate up` | Run database migrations |
| `gator migrate down` | Rollback migrations |

## Development

### Building from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/gator.git
   cd gator
   ```

2. Build the binary:
   ```bash
   go build -o gator .
   ```

### Running Tests

```bash
go test ./...
```

## Deployment

Gator produces a statically compiled binary that can be distributed without dependencies:

```bash
# Build for current platform
go build -o gator .

# Cross-compile for Linux
GOOS=linux GOARCH=amd64 go build -o gator-linux-amd64 .
```

## GitHub Repository

The project is hosted at:  
https://github.com/adibbelel/gator


This README provides users with clear installation instructions, configuration guidance, and usage examples. The structure follows common open-source project conventions while being specific to this CLI tool's functionality.
