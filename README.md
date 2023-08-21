<div align="center">
    <h1><code>🌐</code> Golang MVC</h1>
    <strong>Simple MVC pattern for Golang</strong>
</div>

<br />

## `🧪` Test

[![codecov](https://codecov.io/gh/PunGrumpy/golang-mvc-simple/graph/badge.svg?token=5JWUON145V)](https://codecov.io/gh/PunGrumpy/golang-mvc-simple)
[![🐰 Go](https://github.com/PunGrumpy/golang-mvc-simple/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/PunGrumpy/golang-mvc-simple/actions/workflows/go.yml)

## `📝` About

This is a simple MVC pattern for Golang. It's not a framework, it's just a simple pattern to help you organize your code.

## `📚` How to use

### `📦` Install

```bash
go run cmd/main.go
```

### `📌` Add new soldier

```bash
curl -X POST -H "Content-Type: application/json" -d '{
      "id": 2,
      "name": "Alice",
      "rank": "Sergeant",
      "wife": "Eve",
      "salary": 40000,
      "home": true,
      "car": false,
      "corruption": false
  }' http://localhost:8080/soldier/
```

### `📌` Update soldier infomation

```bash
curl -X PUT -H "Content-Type: application/json" -d '{
      "corruption": true
  }' http://localhost:8080/soldier/2
```

### `📌` Get soldier infomation

```bash
curl http://localhost:8080/soldier/2
```

### `📌` Delete soldier

```bash
curl -X DELETE http://localhost:8080/soldier/2
```

## `📜` License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
