# PayoutManagementSystem

This project is about the payout management system built using golang.

# Project Setup

## Clone the repository

Clone this repo: <a href = "https://github.com/Swarathmica-infraspec/payoutManagementSystem"> source link  </a>

# Requirements

GO-VERSION: 1.22.2 and above

The project contains payoutmanagementsystem/ <br>

    <t> - .github/workflows/payoutManagementSystem.yml <br>
    <t> - payoutmanagementsystem/
    <t> - payee.go <br>
    <t> - payee_test.go <br>
    <t> - go.mod <br>
    <t> - go.sum <br>
    <t> - main.go <br>
    <t> - main_test.go <br>
    <t> - README.md <br>

NOTE: Only email ids with .com are supported.

# Run Tests

Test can be run by executing the below command in the terminal
  go test -v ./...

NOTE: this project is still under development and hence does not have HTTP API now.

# CI

The workflow is triggered on every push and pull request.
It runs the following checks automatically:
- Linting with `golangci-lint`
- Tests with `go test`
