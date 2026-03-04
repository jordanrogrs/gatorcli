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
To use gatorcli, we'll need to setup a .gatorconfig.json file in your home directory with a link to the PostgreSQL server that will tell the application where to find the database. Run the following code in your terminal. Be sure to change the *USERNAME* in the code to your system's username.
```
bash
cat <<EOF > ~/.gatorconfig.json
{
  "db_url": "postgres://USERNAME:@localhost:5432/gatorcli"
}
EOF
```

If you don't know your *USERNAME*, you can find it by using the following command in the terminal:
```
whoami
```

## Usage

Here is a list of commands you can use with gatorcli:
- register: ```gatorcli register <user>``` - creates a new user and automatically logs them in
- login: ```gatorcli login <user>``` - switches the logged in user 
- reset: ```gatorcli reset``` - *CAUTION* - deletes the entire database 
- users: ```gatorcli users``` - lists all users in the database, displays current user logged in 
- addfeed: ```gatorcli addfeed <name> <url>``` - adds a feed to the database 
- feeds: ```gatorcli feeds``` - lists all feeds in the database
- follow: ```gatorcli follow <url>``` - allows user to follow a feed 
- following: ```gatorcli following``` - lists all feeds logged in user is following 
- unfollow: ```gatorcli unfollow``` - removes feed from user profile 
- browse: ```gatorcli browse [limit]``` - displays the latest posts from a users followed feeds 
- agg: ```gatorcli agg <time_between_requests>``` - starts an endless loop to scrape feeds that a user is following and store all the posts

To get started, you'll need to register a user and add your first feed, then you can use the agg command to scrape it! Gatorcli was designed to have one terminal open and scraping the web in the background while you use another terminal to browse the posts. 