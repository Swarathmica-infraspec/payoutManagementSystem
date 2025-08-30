# PayoutManagementSystem

This project is about the payout management system built using golang.

# Project Setup

## Clone the repository

Clone this repo: <a href = "https://github.com/Swarathmica-infraspec/payoutManagementSystem"> source link  </a>

# Requirements

GO-VERSION: 1.22.2 and above

The project contains payoutmanagementsystem/ <br>
- .github/workflows/payoutManagementSystem.yml <br>
- payee/
  - payee.go <br>
  - payee_test.go <br>
  - payee_db.sql <br>
  - payeeDAO.go <br>
  - payeeDAO_test.go <br>
  - payeeAPI.go <br>
  - payeeApi_test.go <br>
- go.mod <br>
- go.sum <br>
- main.go <br>
- main_test.go <br>
- README.md <br>

NOTE: Only email ids with .com are supported.



# Database Setup

We use PostgreSQL running inside Docker for persistant storage.

## 1. Start Postgres with Docker Compose

From the project root, run:

docker compose up -d db


This will:

Start a container named devcontainer-db-1 (from .devcontainer/docker-compose.yml)


## 2. Create Payees Table

Copy the SQL file into the container:

docker cp payee/payee_db.sql devcontainer-db-1:/payee_db.sql


Then apply it:

docker exec -it devcontainer-db-1 psql -U postgres -d postgres -f /payee_db.sql


# Data Access Object

1. payeeDAO contains database query for payee and payeeDAO_test contains relevant tests

To run tests:

docker exec -it devcontainer-app-1 bash

cd /workspaces/payoutManagementSystem

go test -v ./...


# HTTP API Usage

since postgres is run from docker, 

docker exec -it devcontainer-app-1 bash

cd /workspaces/payoutManagementSystem

then run: go run main.go #entry point

payeeApi.go has the code for API while payeeAPI_test.go has test code

NOTE: Supports only POST request

1. POST request 
curl -X POST http://localhost:8080/payees \
  -H "Content-Type: application/json" \
  -d '{
    "name":"Abc",
    "code":"123",
    "account_number":123456789,
    "ifsc":"CBIN012345",
    "bank":"CBI",
    "email":"abc@example.com",
    "mobile":9876543210,
    "category":"Employee"
  }'

expected response: {'id':1}

2. for test: go test ./...  #run inside docker env as above

# Run Tests

Test can be run by executing the below command in the terminal
  go test -v ./...

NOTE: this project is still under development and hence does not have HTTP API now.

# CI

The workflow is triggered on every push and pull request.
It runs the following checks automatically:
- Linting with `golangci-lint`
- Tests with `go test`

