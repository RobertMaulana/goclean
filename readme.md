## Description
This is an example of codebase of Clean Architecture in Go (Golang) projects.

Rule of Clean Architecture:

- Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
- Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
- Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
- Independent of Database. You can swap out Oracle or SQL Server, for Mongo or something else. Your business rules are not bound to the database.
- Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

Credit -> **Uncle Bob**

## Folder structure

- cmd: Contain the application-specific code for the executable applications in the project.
- pkg: Contain the ancil ary non-application-specific code used in the project. We’l use it to hold potential y reusable code like validation helpers and the SQL database models for the project.

## Layer structure

- Entity: Contain model / entity for specific domain business.
- Repository: Contain code layer that handle connecting between database and usecase.
- Usecase / service: Contain business logic code.
- Handler: Contain code layer that handle incoming request from client.

## Code architecture

- SOLID principles and implement clean code

## How to run server

**Make Sure you have run the goclean.sql in your mysql**

##### 1. Docker compose
`make run`

##### 2. Go run
`go run cmd/web/*`

## How to run test
`make test`

## Tools used
##### 1. Echo
##### 2. Mockery
##### 3. Testify
##### 4. Goquery
##### 5. Go-sql-driver
##### 6. Viper

## Todo
- [ ] Security
- [ ] Middleware
- [ ] Standard response
- [ ] Centralize config
- [ ] Repository unit test
- [ ] Helpers unit test
