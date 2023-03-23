# Foundries

## Descriptions

This repository contains the source code for the Foundries, a brain for No-Code AI Platform.

## Project Structure
We use the concept of Clean Architecture to structure our project. The project is divided into 4 layers:
- Entities / Models -> This layer contains the data models that are used in the application.
- Use Cases / Services -> This layer contains the business logic of the application.
- Interfaces / Controllers -> This layer contains the controllers that are used to handle the requests.
- Infrastructure / Database -> This layer contains the database connection and the database models.

## Dependencies
- [Gin]()
- [GORM]()


## Installation
Make sure to install all dependencies before running the application.
```
go mod tidy
```
Then run the application.
```
go run main.go
```

## Usage
The application is a REST API that can be accessed through the following endpoints:
```
GET /users
GET /users/:id
POST /users
```

## TODO
Here are some todos that should be done in the future:
- [ ] Add dataset endpoint
- [ ] Add project endpoint
- [ ] Add auth middleware
- [ ] Add unit tests
- [ ] Add integration tests
- [ ] Add CI/CD pipeline
- [ ] Add Dockerfile
- [ ] Add Kubernetes deployment