# Blog-Aggregator
create a config file in your home directory `~/.gatorconfig.json` with this content:
```
{
  "db_url": "postgres://example"
}
```
## requierd packets:
`go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`
`go get github.com/lib/pq`
`go get github.com/google/uuid`

## postgres install (at least version 15):
### linux:
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
you will querie against the database that is set in the config file `~/.gatorconfig.json`, example for the gator database:
```
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
}
```
if you have a custom setup (the protocol has to be postgres!):
```
{
  "db_url": "protocol://username:password@host:port/database?sslmode=disable"
}
```
to remove the service:
```
sudo systemctl disable <service-name>
```