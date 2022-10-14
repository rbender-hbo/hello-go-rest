# Simple Go REST API

Uses `chi` library for URL routing.

Run using `./start.sh`.

Exposes:

- All Foo - http://localhost:8080/foo
- Individual Foo - http://localhost:8080/foo/{fooId}

## TODO

- Handle updating existing Foo objects via PUT
- Use logging framework
- Try dependency injection framework
- Create Makefile
- Database persistence? (Sqlite?)
