# Simple Go REST API

Uses `chi` library for URL routing.

Run using `./start.sh`.

Exposes:
  - All Foo - http://localhost:8080/foo
  - Individual Foo - http://localhost:8080/foo/{fooId}

## TODO

- Handle creating new Foo objects via POST
- Handle updating existing Foo objects via PUT
- Try dependency injection framework
- Create Makefile
- Database persistence? (Sqlite?)
