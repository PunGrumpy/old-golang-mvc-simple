<div align="center">
    <h1><code>🌐</code> Golang MVC</h1>
    <strong>Simple MVC pattern for Golang</strong>
</div>

<br />

## `📝` About

This is a simple MVC pattern for Golang. It's not a framework, it's just a simple pattern to help you organize your code.

## `📚` How to use

### `📦` Install

```bash
go run .
```

### `📝` Post

#### `📌` Add

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

#### `📌` Update

```bash
curl -X PUT -H "Content-Type: application/json" -d '{
      "corruption": true
  }' http://localhost:8080/soldier/2
```

#### `📌` Get

```bash
curl http://localhost:8080/soldier/2
```

#### `📌` Eat Tax'

```bash
curl -X PUT -H "Content-Type: application/json" -d '{
      "salary": 40000
  }' http://localhost:8080/soldier/eat/10000
```

## `📜` License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

```

```
