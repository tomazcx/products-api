# PRODUCTS API

A RESTful API made with Golang that features a CRUD for Products and Users with JWT authentication.

Libraries used:

- Chi, for HTTP routing
- Testify, for automated testing
- GORM, for database management

## Getting started

To run this project on your local machine, follow these steps:

1- Clone the repository

```bash
git clone https://github.com/tomazcx/odonto-dashboard-go.git
```

2- Run the docker container using `docker-compose`

```bash
docker-compose up
```

## Swagger Docs

All the api endpoints are documented with swagger. You can access it at /docs/index.html

## Testing

Run the following command in the root of the project to execute all the application tests:

```bash
go test ./...
```
