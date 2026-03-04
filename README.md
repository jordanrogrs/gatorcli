# gatorcli

A CLI blog aggregator written in Go that allows users to store RSS feeds and posts in a local PostgreSQL database.

## Requirements

Before running gatorcli you will need:
- Go installed, [download + install here](https://go.dev/doc/install)
- PostgreSQL installed, [download + install here](https://www.postgresql.org/download/)

## Installation

### Step 1: Install
To install gatorcli on your system, go to your preferred terminal application and run:
```
go install github.com/jordanrogrs/gatorcli@latest
```

### Step 2: Config Setup
To use gatorcli, we'll need to set up a .gatorconfig.json file in your home directory with a link pointing to the PostgreSQL database used by gatorcli. Run the following command in your terminal. Be sure to change the ```<username>``` to your system's username.
```bash
cat <<EOF > ~/.gatorconfig.json
{
  "db_url": "postgres://<username>@localhost:5432/gatorcli"
}
EOF
```

If you don't know your system username, you can find it by using the following command in the terminal:
```bash
whoami
```

After installation, the ```gatorcli``` command will be available in your terminal.

## Quick Start

* Register a user:
```bash
gatorcli register jordan
```

* Add a feed:
```bash
gatorcli addfeed "Hacker News" "https://news.ycombinator.com/rss"
```

* Follow the feed:
```bash
gatorcli follow "https://news.ycombinator.com/rss"
```

* Start the aggregator:
```bash
gatorcli agg 1m
```

* Browse posts:
```bash
gatorcli browse
```


## Usage

Available gatorcli commands:
#### User Commands
- **register**: ```gatorcli register <user>``` - creates a new user and automatically logs them in
- **login**: ```gatorcli login <user>``` - switches the logged in user 
- **users**: ```gatorcli users``` - lists all users in the database, displays current user logged in
#### Feed Commands
- **addfeed**: ```gatorcli addfeed <name> <url>``` - adds a feed to the database 
- **feeds**: ```gatorcli feeds``` - lists all feeds in the database
- **follow**: ```gatorcli follow <url>``` - allows user to follow a feed 
- **following**: ```gatorcli following``` - lists all feeds logged in user is following 
- **unfollow**: ```gatorcli unfollow``` - removes feed from user profile
#### Content Commands
- **browse**: ```gatorcli browse [limit]``` - displays the latest posts from a user's followed feeds, default is 2 posts
- **agg**: ```gatorcli agg <time_between_requests>``` - continuously scrapes the logged in user's feeds at a specified interval
#### Utility Commands
- **reset**: ```gatorcli reset``` - **CAUTION** - deletes the entire database 

To get started, you'll need to register a user and add your first feed, then you can use the agg command to scrape it! Gatorcli was designed to have one terminal open and scraping the web in the background while you use another terminal to browse the posts. 