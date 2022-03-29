# littSQL

## Get started

#### 1. Setup the database

If PostgreSQL is not installed on your system, make sure to install it from [here](https://www.postgresql.org/download/) and make sure it runs.

Now run the [init.sql](./database/init.sql) file inside of your PostgreSQL database, to create all needed schemas, tables, and stored procedures.

#### 2. Setup backend

First make sure that the go cli is installed, by running `where go`, if it tells you that go has not been found, install go from [here](https://go.dev/dl/).

In the backend folder, now run `go get` to install the go dependecies needed for the project.

After that create a .env file in the [backend](./backend/) directory, to configure the backend with the PostgreSQL credentials. Below is a sample .env with important key-value pairs and their respective defaults:

```bash
PG_HOST="127.0.0.1"
PG_NAME="postgres"
PG_USER=""
PG_PASSWORD="postgres"
```

To use the default of a key, you can omit that key-value pair. If PostgreSQL has been installed without much configuration, you probably only need to configure the `PGUSER` key with your home directory name.

Now in the [backend](./backend/) directory run `go run test_db/test_db.go` and watch the console for errors if postgres has not been configured correctly, to adjust configurations and rerun until it works.

#### 3. Setup frontend

First make sure that the nodejs cli is installed, by running `where npm`, if it tells you that npm has not been found, install nodejs from [here](https://nodejs.org/en/download/).

In the backend folder, now run `go get` to install the go dependecies needed for the project.

To install the dependencies needed, run `npm install` in the [frontend](./frontend/) directory.

#### 4. Run the project (development)

First run the backend project from the [backend](./backend/) directory with `go run main.go`.

Then run the frontend project from the [frontend](./frontend/) directory with `npm run dev`.

#### 5. More configurations

#### 6. Run the project (production)
