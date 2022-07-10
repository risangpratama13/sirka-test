# Sirka Assessment Study Case

## Live Demo
https://sirka-test.herokuapp.com/MyWeb/

## Description
This project has 4 domain layer: 
- Model Layer
- Repository Layer
- Service Layer
- Controller Layer

## How to Run This Project
### Prerequisites:
- Golang 1.18
- Docker Compose

### Run Testing:
```bash
$ go test -v ./...
```

### Run With docker-compose:
```bash
# Move to directory
$ cd workspace

# Clone into YOUR $GOPATH/src
$ git clone https://github.com/risangpratama13/sirka-test

# Move to project
$ cd sirka-test

# Run docker-compose
$ docker-compose up -d

# check if the containers are running (optional)
$ docker ps
```
Open browser or postman and access url http://localhost:8001/

**NB: by default the app will run on port 8001, but you can change it in docker-compose.yml**

### Run The Project:
```bash
# Move to directory
$ cd workspace

# Clone into YOUR $GOPATH/src
$ git clone https://github.com/risangpratama13/sirka-test

# Move to project
$ cd sirka-test

# Run docker-compose
$ docker-compose up -d

# check if the containers are running (optional)
$ docker ps

# Download dependencies
$ go mod download

# Run App
$ go run main.go
```
Open browser or postman and access url http://localhost:8080/

### Notes
- By default when the application runs the table will be migrated automatically
- Users data will be empty, please import the data from the `test_case_users.sql` file

## Tools and Libraries Used:
- Gin Framework
- Viper
- Gorm
- Testify
- SQLMock