# Go Aquila DB

Copy .env_dist to .env
```bash
$ cp .env_dist .env
```

## Run app in docker.

1. Build: --- only for the first time!!!
```bash
$ docker-compose build
```
Then if need change your credentials in .env file. (No need by default)

2. Run app:
```bash
$ docker-compose up
```

Run migrations from migrations/migration.sql file:
```bash
$ MIGRATIONS=`cat migrations/migration.sql` && DATABASE_NAME=$(grep DB_NAME .env | cut -d '=' -f 2-) && docker exec -it postgres_db psql "postgresql://$DB_USER:$DB_PASS/${DATABASE_NAME}" --command="$MIGRATIONS"
```

Run user seeder from migrations/seeder.sql file:
```bash
$ SEEDER=`cat migrations/seeder.sql` && DATABASE_NAME=$(grep DB_NAME .env | cut -d '=' -f 2-) && docker exec -it postgres_db psql "postgresql://$DB_USER:$DB_PASS/${DATABASE_NAME}" --command="$SEEDER"
```
Every user credentials you can find in migrations/seeder.sql file. Every users have password = 123
Admin - admin@admin.com 


3. To stop - press Ctrl+c and then:
```bash
$ docker-compose down
```

## Notes.

All credentials you can find in .env file.

1. Server working on the port :8000
2. To see database in the browser \
    localhost:8085  \
    Then login with credentials:  \
    ```
        Database - PostgreSQL
        Server - db
        User - root
        Password - password
        Db_name - go_aquila
    ```
    