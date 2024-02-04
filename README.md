# PostgreSQL jsonb type support for go
![go workflow](https://github.com/asif-mahmud/pg-jsonb/actions/workflows/go.yml/badge.svg)

This is a small package to support [PostgreSQL jsonb](https://www.postgresql.org/docs/current/datatype-json.html) data type.

## Features
1. Fully tested.
2. Supports PostgreSQL jsonb type by sql Scanner and Valuer interfaces
3. Supports go json Marshaller and Unmarshaller interfaces
4. Supports go Stringer for json printing

## Installation
To add the package to your project run -

```
go get -u github.com/asif-mahmud/pg-jsonb
```

### Documentation

godoc: [https://pkg.go.dev/github.com/asif-mahmud/pg-jsonb](https://pkg.go.dev/github.com/asif-mahmud/pg-jsonb)

## Version history

### Version 0.8.0

Initial version with tests for scanner and valuer interfaces.

