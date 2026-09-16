# 🐊 Gator

A command-line RSS feed aggregator written in **Go** with **PostgreSQL**.

Gator lets you register users, add and follow RSS feeds, periodically fetch new posts, and browse aggregated content.

---

## 📋 Requirements

Before getting started, you'll need:

* **[Go](https://go.dev/)** — version 1.21 or later recommended
* **[PostgreSQL](https://www.postgresql.org/)** — used to store users, feeds, and posts

---

## 🚀 Installation

Clone the repository and navigate into the project:

```bash
git clone <your-repository-url>
cd blog-aggregator
```

Install the `gator` CLI with Go:

```bash
go install .
```

This installs the `gator` executable into Go's binary directory.

If your shell cannot find the `gator` command, make sure Go's binary directory is included in your `PATH`:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

You can verify the installation with:

```bash
gator
```

---

## 🗄️ Database Setup

Gator requires a PostgreSQL database.

Create a database for Gator:

```bash
createdb gator
```

You'll need your PostgreSQL connection string in the following format:

```text
postgres://username:password@localhost:5432/gator?sslmode=disable
```

Replace the username, password, and database name with your own PostgreSQL configuration.

---

## ⚙️ Configuration

Gator uses a configuration file located in your home directory:

```text
~/.gatorconfig.json
```

Create the file:

```bash
touch ~/.gatorconfig.json
```

Then add your PostgreSQL connection string:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable"
}
```

For example:

```json
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
}
```

> **Note:** Make sure the database exists and that the credentials in `db_url` are valid before running Gator.

---

## 🎮 Using Gator

### 👤 Register a user

Create a new Gator user:

```bash
gator register <username>
```

Example:

```bash
gator register gerard
```

### 🔑 Log in

Set the current user:

```bash
gator login <username>
```

Example:

```bash
gator login gerard
```

### 📡 Add a feed

Add an RSS feed:

```bash
gator addfeed <feed-name> <feed-url>
```

Example:

```bash
gator addfeed "The Joe Rogan Experience" https://feeds.megaphone.fm/GLT1412515089
```

Adding a feed also follows it for the currently logged-in user.

### 📰 List feeds

View the feeds that have been added:

```bash
gator feeds
```

### ➕ Follow a feed

Follow an existing feed:

```bash
gator follow <feed-url>
```

Example:

```bash
gator follow https://feeds.megaphone.fm/GLT1412515089
```

### 👀 View followed feeds

See the feeds followed by the current user:

```bash
gator following
```

### 🔄 Aggregate posts

Start periodically fetching posts from RSS feeds:

```bash
gator agg <duration>
```

For example:

```bash
gator agg 1m
```

This fetches feeds and waits one minute between requests.

Other valid durations include:

```text
10s
30s
5m
1h
```

### 📖 Browse posts

View recently aggregated posts:

```bash
gator browse
```

---

## 🏁 Quick Start

Once PostgreSQL and Gator are configured, you can get started with:

```bash
# Register
gator register gerard

# Log in
gator login gerard

# Add a feed
gator addfeed "The Joe Rogan Experience" https://feeds.megaphone.fm/GLT1412515089

# Start following the feed
gator follow https://feeds.megaphone.fm/GLT1412515089

# Start aggregating posts
gator agg 1m
```

Then, in another terminal:

```bash
gator browse
```

---

## 📚 Command Reference

| Command                | Purpose                                |
| :--------------------- | :------------------------------------- |
| `register <username>`  | Register a new user                    |
| `login <username>`     | Log in as an existing user             |
| `addfeed <name> <url>` | Add an RSS feed                        |
| `feeds`                | List available feeds                   |
| `follow <url>`         | Follow an existing feed                |
| `following`            | List the current user's followed feeds |
| `agg <duration>`       | Periodically fetch RSS feeds           |
| `browse`               | Browse aggregated posts                |

---

<p align="center">
  Built with Go
</p>
