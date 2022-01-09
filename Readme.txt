Init project;
    $ go mod init aquiladb

Install gin:
    $ go get -u github.com/gin-gonic/gin

Install dependencies:
    $ go get github.com/lib/pq
    $ go get github.com/jmoiron/sqlx
    $ go get github.com/golobby/dotenv
    $ go get -u github.com/golang-jwt/jwt

// ======================================

Run app:
    $ go run cmd/main.go


Run with live reload:
    Install air:
        $ go get -u github.com/cosmtrek/air
    Init air into the project:
        $ air init  
        or
        $ ~/go/bin/air init
        It will create the file ".air.toml" then put settings inside the file.

    Run:
        $ ~/GOPATH/bin/air
        or 
        $ ~/go/bin/air

// ======================================

Postman.
-----------------

1. Create temp customer:
    post /customer

2. Create permanent customer:
    auth token - required
    patch /customer
    {
        "first_name": "Bob",
        "last_name": "Stone",
        "secret_key": "79hNcex4fyfQLcO",
        "email": "bob@bob.com"
    }

3. Get customer:
    auth token - required
    get /customer

4. Auth:
    post 
    {
        "secret_key": "tUJqvSphQiPLnUp"
    }
    token in response


// ======================================

// ======================================
// ======================================
// ======================================
// ======================================


Example for random name:
https://github.com/moby/moby/blob/master/pkg/namesgenerator/names-generator.go


Run user seeder from migrations/seeder.sql file:
```bash
$ SEEDER=`cat migrations/seeder.sql` && DATABASE_NAME=$(grep DB_NAME .env | cut -d '=' -f 2-) && docker exec -it postgres_db psql "postgresql://$DB_USER:$DB_PASS/${DATABASE_NAME}" --command="$SEEDER"
```