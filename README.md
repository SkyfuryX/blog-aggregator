# Blog Aggregator

A command-line RSS feed aggregator built with Go and PostgreSQL.

Repository: https://github.com/SkyfuryX/blog-aggregator

---

# Installation

## Prerequisites

Before installing the project, make sure you have the following installed:

- [Go](https://go.dev/doc/install)
- [PostgreSQL](https://www.postgresql.org/download/)

You will also need a running PostgreSQL database instance.

---

## 1. Clone the Repository

```bash
git clone https://github.com/SkyfuryX/blog-aggregator.git
cd blog-aggregator
```

---

## 2. Create a PostgreSQL Database

Open the PostgreSQL shell:

```bash
psql postgres
```

Create a database:

```sql
CREATE DATABASE blog_aggregator;
```

---

## 3. Install the Application

Install the CLI globally using Go:

```bash
go install github.com/SkyfuryX/blog-aggregator@latest
```

If you prefer to build locally:

```bash
go build
```

---

## 4. Configure the Application

Create a configuration file in your home directory named:

```text
.gatorconfig.json
```

Example configuration:

```json
{
  "db_url": "postgres://username:password@localhost:5432/blog_aggregator?sslmode=disable"
}
```

Replace:

- `username` with your PostgreSQL username
- `password` with your PostgreSQL password
- `blog_aggregator` with your database name

---

# Usage

## Register a User

```bash
blog-aggregator register <username>
```

Example:

```bash
blog-aggregator register gannon
```

---

## Log In

```bash
blog-aggregator login <username>
```

---

## Add an RSS Feed

```bash
blog-aggregator addfeed <feed_url>
```

Example:

```bash
blog-aggregator addfeed https://news.ycombinator.com/rss
```

---

## View All Feeds

```bash
blog-aggregator feeds
```

---

## Follow an Existing Feed

```bash
blog-aggregator follow <feed_url>
```

---

## Unfollow a Feed

```bash
blog-aggregator unfollow <feed_url>
```

---

## Start the Feed Aggregator

Run the scraper at a specified interval:

```bash
blog-aggregator agg 30s
```

Example intervals:

```bash
blog-aggregator agg 10s
blog-aggregator agg 1m
blog-aggregator agg 5m
```

This continuously fetches and stores posts from subscribed RSS feeds.

---

## Browse Posts

```bash
blog-aggregator browse
```

Limit the number of posts returned:

```bash
blog-aggregator browse 10
```

---

## List Users

```bash
blog-aggregator users
```

---

# Example Workflow

```bash
# Register a user
blog-aggregator register gannon

# Log in
blog-aggregator login gannon

# Add a feed
blog-aggregator addfeed https://news.ycombinator.com/rss

# Start scraping feeds every 30 seconds
blog-aggregator agg 30s

# Browse collected posts
blog-aggregator browse 10
```

---

# Features

- RSS feed aggregation
- PostgreSQL-backed storage
- Multi-user support
- Feed following and unfollowing
- Periodic feed scraping
- Terminal-based browsing experience