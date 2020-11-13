# REST API written in GO

Simple API project uses Gorilla toolkit. The application has one endpoint with a JSON formatted response.

## Install dependencies
```
go get .
```

## Start application
```
go run .
```

## Example request and response
```
GET http://localhost:8000/users/all
---

Response:
{
    "data": [
        "Peter",
        "David",
        "Dora"
    ]
}
```
