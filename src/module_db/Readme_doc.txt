


// ==============================================


#############################################3
    AquilaDB
#############################################3


1. Create database:
--------------------

http://localhost:5001/db/create

Request header:
Content-Type - application/json

Request body:
{
    "data": {
        "schema": {
            "description": "this is my database",
            "unique": "r8and0mseEd901",
            "encoder": "strn:msmarco-distilbert-base-tas-b",
            "codelen": 768,
            "metadata": {
                "name": "string",
                "age": "number"
            }
        }
    },
    "signature": "secret"
}

{
  "data": {
    "schema": {
      "codelen": 768, 
      "description": "Database of Bold", 
      "encoder": "strn:msmarco-distilbert-base-tas-b", 
      "metadata": {
        "age": "number", 
        "name": "Bold"
      }, 
      "unique": "y3fpVc7wEk0bOU"
    }
  }, 
  "signature": "secret"
}

Response:
{
    "database_name": "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
    "success": true
}

// ==============================================

http://localhost:5001/db/doc/insert

post
{
   "data": {
       "docs": [
           {
           "payload":
               {
                   "metadata": {
                       "name":"name1",
                       "age": 20
                   },
                   "code": [0.1, 0.2, 0.3]
               }
           },
           {
           "payload":
               {
                   "metadata": {
                       "name":"name2",
                       "age": 30
                   },
                   "code": [0.4, 0.5, 0.6]
               }
           }
       ],
       "database_name": "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7"
   },
   "signature": "secret"
}


response
{
    "ids": [
        "3gwTnetiYJfHTBcqGwoxETLsmmdGYVsd5MRBohuTG22C",
        "BXsbHy9B3tU9zaHwU41jATzDBisNEFa67XKvYZhB2fzQ"
    ],
    "success": true
}

// ==============================================

http://localhost:5001/db/doc/delete
request:
{ 
    "data": {
        "ids": [
            "3gwTnetiYJfHTBcqGwoxETLsmmdGYVsd5MRBohuTG22C",
            "BXsbHy9B3tU9zaHwU41jATzDBisNEFa67XKvYZhB2fzQ"
        ], 
        "database_name": "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7"
    },
    "signature": "secret"
}

{
  "data": {
    "database_name": "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7", 
    "ids": [
      "3gwTnetiYJfHTBcqGwoxETLsmmdGYVsd5MRBohuTG22C", 
      "BXsbHy9B3tU9zaHwU41jATzDBisNEFa67XKvYZhB2fzQ"
    ]
  }, 
  "signature": "secret"
}

response:
{
    "ids": [
        "3gwTnetiYJfHTBcqGwoxETLsmmdGYVsd5MRBohuTG22C",
        "BXsbHy9B3tU9zaHwU41jATzDBisNEFa67XKvYZhB2fzQ"
    ],
    "success": true
}


// ==============================================


Stop container:
    $ docker-compose -p "aquilanet" down

// ==============================================
// ==============================================

Make sure to create database like in the first step:
    http://localhost:5002/prepare
    Request header:
    Content-Type - application/json

    Request body:
    {
        "data": {
            "schema": {
                "description": "this is my database",
                "unique": "r8and0mseEd901",
                "encoder": "strn:msmarco-distilbert-base-tas-b",
                "codelen": 768,
                "metadata": {
                    "name": "string",
                    "age": "number"
                }
            }
        },
        "signature": "secret"
    }

1. html from the browser extract by browser extension.
2. Browser extension send html string to go project.

3. Go project should take html string.

4. Then send html string to external service -> mercury 5009
    header
        Content-Type: application/json
    body
        POST
        http://localhost:5009/process
        {
            "url": "http://test.com",
            "html": "<!DOCTYPE html><html><head><title>Bla</title></head><body><h1>Test Aqula DB</h1><p>At the time, no single team member knew Go, but within a month, everyone was writing in Go and we were building out the endpoints. It was the flexibility, how easy it was to use, and the really cool concept behind Go (how Go handles native concurrency, garbage collection, and of course safety+speed.) that helped engage us during the build. Also, who can beat that cute mascot!</p></body></html>"
        }


4. Ater getting response extract content.
5. Send content to txtpick
    header
        Content-Type: application/json
    body
    http://localhost:5008/process
        POST
        {   
            "url": "http://test.com",
            "html": "<body><p>At the time, no single team member knew Go, but within a month, everyone was writing in Go and we were building out the endpoints. It was the flexibility, how easy it was to use, and the really cool concept behind Go (how Go handles native concurrency, garbage collection, and of course safety+speed.) that helped engage us during the build. Also, who can beat that cute mascot!</p></body>"
        }

    response 
    {
        "result": [
            "It was the flexibility, how easy it was to use, and the really cool concept behind Go (how Go handles native concurrency, garbage collection, and of course safety+speed.)"
        ],
        "success": true
    }

6. After getting response. Take array from result.
7. Sent an array with help go_module to AquilaHub
    http://localhost:5002/compress
    header
        Content-Type: application/json
    body post:
        {
            "data": {
                "text": [
                    "It was the flexibility, how easy it was to use, and the really cool concept behind Go (how Go handles native concurrency, garbage collection, and of course safety+speed.)"
                ],
                "databaseName": "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7"
            }
        }

    response will be vectors array of arrays.
    Loop through it and make single payload in single array.

8. Take vectors and make 
    http://localhost:5001/db/doc/insert

    with tokek vectors
    every array is single object in payload