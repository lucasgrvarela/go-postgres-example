# Go Postgres Example

This is an example project that demonstrates the use of the Repository Pattern with PostgreSQL in Go.

## Prerequisites

- Go programming language
- Docker (optional, for running PostgreSQL in a container)

## Getting Started

1. Clone the repository to your local machine.

2. Set up a PostgreSQL database. You can either install PostgreSQL locally or use a Docker container.

- If you have PostgreSQL installed locally, make sure it's running and create a new database.

- If you want to use a Docker container, run the following command to start a PostgreSQL container:

```bash
docker run --name my-postgres-container -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres
```

3. Set up the database schema by applying the migration scripts located in the `migrations` folder.

- Ensure that you have the `migrate` tool installed. You can install it following the tutorial [here](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate).

- Apply the migrations using the following command:

```bash
migrate -database "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable" -path ./migrations up
Output: 1/u create_users_table (14.613453ms)
```
Optional: you can validate the migration really created the tables inside the docker container postgres instance.
```
$ migrate -path migrations/ -database "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable" up
1/u create_users_table (14.613453ms)

$ docker exec -it my-postgres-container bash
$ psql -U postgres -d postgres # run inside the container
$ postgres=# \dt
               List of relations
 Schema |       Name        | Type  |  Owner
--------+-------------------+-------+----------
 public | schema_migrations | table | postgres
 public | users             | table | postgres
(2 rows)

$ postgres=# \d+ users
                                                                Table "public.users"
  Column  |          Type          | Collation | Nullable |              Default              | Storage  | Compression | Stats target | Description
----------+------------------------+-----------+----------+-----------------------------------+----------+-------------+--------------+-------------
 id       | integer                |           | not null | nextval('users_id_seq'::regclass) | plain    |             |              |
 username | character varying(50)  |           | not null |                                   | extended |             |              |
 email    | character varying(100) |           | not null |                                   | extended |             |              |
Indexes:
    "users_pkey" PRIMARY KEY, btree (id)
```

4. Update the `connString` variable in the `main.go` file with your PostgreSQL connection details.

```go
connString := "host=localhost port=5432 user=postgres password=mysecretpassword dbname=postgres sslmode=disable"
```

5. Run the project:
```bash
go run main.go
```

## Project Structure
```
.
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── migrations
│   ├── 1_create_users_table.down.sql
│   └── 1_create_users_table.up.sql
├── README.md
└── user
    ├── user.go
    └── user_repository
        ├── postgres_user_repository.go
        └── user_repository.go
```
* main.go: Contains the main entry point of the application.
* migrations: Directory containing the database migration scripts.
* user: Directory for the user domain.
    * user.go: Definition of the User struct.
    * user_repository: Directory for the user repository.
        * user_repository.go: Definition of the UserRepository interface.
        * postgres_user_repository.go: Implementation of the UserRepository interface using PostgreSQL.
