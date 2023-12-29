# ps-tag-onboarding-go

## The project
This is the Go version of the [onboarding exercise](https://github.com/wexinc/ps-tag-onboarding) for the Tag Team.

## Frameworks, libs and technologies
 - [Go](https://go.dev/)
 - [Gin](https://github.com/gin-gonic/gin)
 - [Viper](https://github.com/spf13/viper)
 - [MongoDB](https://www.mongodb.com/)
 - [Docker](https://www.docker.com/)

## Pre-requisites
- Docker
- Go

## Getting started
To run the application, make sure you have Docker and Go installed.
1. Clone the repository
```shell
git clone https://github.com/lucasolsi-wex/ps-tag-onboarding-go
```

2. Navigate to ps-tag-onboarding-go
```shell
cd ps-tag-onboarding-go
```

3. Switch to development branch:
```shell
git checkout development
```

4. Run the database container:
```shell
docker-compose up -d
```

5. Compile the application
```shell
go run .
```

## Endpoints
- **POST** /user
```shell
curl --request POST \
--url http://localhost:8080/user \
--header 'Content-Type: application/json' \
--data '{
"firstName": "First",
"lastName": "Name",
"email": "email@email.com",
"age": 19
}'
```

- **GET** /user/{id}
```shell
curl --request GET \
  --url http://localhost:8080/user/65804b360a2378fc9b806899
```

