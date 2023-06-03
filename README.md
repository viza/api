# Repository for learning REST API development mainly with Go lang backend
# Codebase for some courses (Udemy etc)

## Simple API example
/api/udemy/rest-based-microservices-api-development-in-go-lang/simple

## Examples with customers retrieval by their status (active / inactive)
viza version: /api/udemy/rest-based-microservices-api-development-in-go-lang/customers-by-status-viza-version/
Udemy version: /api/udemy/rest-based-microservices-api-development-in-go-lang/customers-by-status/

## Example with logger
/api/udemy/rest-based-microservices-api-development-in-go-lang/customers-logger/
### Logger
zap logger is used
```
go get -u go.uber.org/zap
go: downloading go.uber.org/zap v1.24.0
go: downloading go.uber.org/multierr v1.11.0
go: downloading go.uber.org/atomic v1.11.0
go: added go.uber.org/atomic v1.11.0
go: added go.uber.org/multierr v1.11.0
go: added go.uber.org/zap v1.24.0
```

## Example with Database sqlx package
/api/udemy/rest-based-microservices-api-development-in-go-lang/customers-sqlx
```
go get github.com/jmoiron/sqlx   
go: downloading github.com/jmoiron/sqlx v1.3.5
go: added github.com/jmoiron/sqlx v1.3.5
```

## Script for database creation and test data
/api/udemy/rest-based-microservices-api-development-in-go-lang/db/resources/database.sql
