# Simple Go REST API

Uses `chi` library for URL routing.

Run using `make all` (executes tests first) or `make run`.

Exposes:

- Get All Foo - GET http://localhost:8080/foo
- Create New Foo - POST http://localhost:8080/foo
- Get Individual Foo - GET http://localhost:8080/foo/{fooId}
- Update Existing Foo - PUT http://localhost:8080/foo/{fooId}

## TODO

- Try dependency injection framework
- Database persistence? (Sqlite?)
