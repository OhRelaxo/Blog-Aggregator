# Blog-Aggregator
create a config file in your home directory `~/.gatorconfig.json` with this content:
```
{
  "db_url": "postgres://example"
}
```

## install go
### on Linux (debian/ubuntu):
```
sudo apt install go
```
### on macOS with brew:
´´´
brew install go
´´´
## postgres install (you need at least version 15):
### Linux:
install it with your paket manager:
```
sudo apt update
sudo apt install postgresql postgresql-contrib
```
make sure its installt and has the right version:
```
psql --version
```
set a password:
```
sudo passwd postgres
```
start the server:
```
sudo service postgresql start
```
or
```
sudo systemctl start postgresql
```

### macOS with brew:
install:
```
brew install postgresql@15
```
make sure the installed version is at least 15:
```
psql --version
```
start the server:
on Linux: 
```
sudo service postgresql start
```
on macOS:
```
brew services start postgresql@15
```
## postgres setup
enter the psql shell:
macOS: ```psql postgres```
Linux: ```sudo -u postgres psql```

your command prompt should look like this:
```
postgres=#
```
create a new database:
```
CREATE DATABASE gator;
```
connect to the new database:
```
\c gator
```
your command prompt should look like this:
```
gator=#
```
if your on Linux set the user password (not on macOS):
```
ALTER USER postgres PASSWORD 'postgres';
```
to exit use
```
exit
```
now you need to create the `~/.gatorconfig.json` in your home dircetory (you can just use nano or vim and then copy one of the following configurations in)
config file this file determines the database you will query against and your user:
```
{
  "db_url": "protocol://username:password@host:port/database?sslmode=disable",
  "current_user_name": "username_goes_here"
}
```
if you just followed the steps above, and you're on Linux this should work for you:
```
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "current_user_name": "username_goes_here"
}
```
if you are on macOS this should work for you:
```
{
  "db_url": "postgres://postgres:@localhost:5432/gator?sslmode=disable",
  "current_user_name": "username_goes_here"
}
```
if you want to remove the postgres service (if you do this you cant use the programm anymore):
```
sudo systemctl disable postgresql
```

## install the program
```
go install github.com/OhRelaxo/Blog-Aggregator
```

## some commands
to start you have to register yourself
```
register <name>
```
after that you need to login
```
login <name>
```
now you can subscribe to feeds
```
addfeed <feed name> <url>
```
with this command you can fetch the feeds
```
agg <time e.g. 10s>
```
with this command you can browse your fetched feeds
```
browse
```