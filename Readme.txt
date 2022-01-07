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

1. Register user:
    post /auth/register
    {
        "first_name": "Bob",
        "last_name": "Bobin",
        "email": "bob@bob.com",
        "password": "123"
    }
    token will be in return

2. Login user:
    post /auth/login
    {
        "email": "bob@bob.com",
        "password": "123"
    }
    token will be in return

3. Protected route:
    get /api/secret
    you should use jwt token

4. Admin route:
    get /admin 
    only for 
    
// ======================================

1. Create temp customer:
    post /customer



// ======================================

// ======================================
// ======================================
// ======================================
// ======================================



