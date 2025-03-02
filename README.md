# Gator - RSS Feed Aggregator

A CLI tool for aggregating and managing RSS feeds with a PostgreSQL database.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Database Setup](#database-setup)
- [Configuration](#configuration)
- [Usage](#usage)
  - [Register New User](#register-new-user)
  - [Login](#login)
  - [Show Users](#show-users)
  - [Add RSS Feed](#add-rss-feed)
  - [Show Feeds](#show-feeds)
  - [Aggregate](#aggregate)
  - [Follow a Feed](#follow-a-feed)
  - [Unfollow a Feed](#unfollow-a-feed)
  - [Show Feeds Followed](#show-feeds-followed)
  - [Show Posts](#show-posts)

## Prerequisites

- [Go](https://go.dev/doc/install) (version 1.24.0 or higher)
- [PostgreSQL](https://www.postgresql.org/download) (version 17.4 or higher)

Alternatively, you can use [devbox](https://www.jetpack.io/devbox) for development:

```sh
devbox shell

# Start PostgreSQL service
devbox services start postgresql
```

## Installation

```sh
go install github.com/rhafaelc/blog-aggregator/cmd/gator
```

## Database Setup

1. Initialize a new PostgreSQL database cluster (if needed):

```sh
initdb
```

2. Create a new PostgreSQL user:

```sh
createuser --interactive
```

3. Create a database for Gator:

```sh
createdb gator
```

4. Run migration:

```sh
# Install goose for migrations
go install github.com/pressly/goose/v3/cmd/goose@latest

# Copy the environment example file
cp .env.example .env

# Run migrations (located in sql/schema)
goose up
```

## Configuration

Create a configuration file at `~/.gatorconfig.json`:

```sh
echo '{"db_url":"postgres://postgres:@localhost:5432/gator?sslmode=disable"}' > ~/.gatorconfig.json
```

## Usage

### Register New User

Register a new user with the specified username.

```sh
gator register <username>
```

### Login

Login with the specified username.

```sh
gator login <username>
```

### Show Users

Display a list of all registered users.

```sh
gator users
```

### Add RSS Feed

Add a new RSS feed with the specified title and URL.

```sh
gator addfeed <title> <url>
```

### Show Feeds

Display a list of all RSS feeds.

```sh
gator feeds
```

### Aggregate

Aggregate RSS feeds with the specified time interval between requests.

```sh
gator agg <time_between_request>
```

### Follow a Feed

Follow an RSS feed with the specified URL.

```sh
gator follow <url>
```

### Unfollow a Feed

Unfollow an RSS feed with the specified URL.

```sh
gator unfollow <url>
```

### Show Feeds Followed

Display a list of all followed RSS feeds.

```sh
gator following
```

### Show Posts

Display a list of posts with an optional limit (default: 2).

```sh
gator browse <limit (default: 2)>
```

```

```
