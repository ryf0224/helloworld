# helloworld

A stateless HTTP web service written in Go. Returns `helloworld` when the client sends `hi`, otherwise returns an error.

## Usage

### Run

```bash
CGO_ENABLED=0 go run .
```

### Request

```bash
curl "localhost:8080/?msg=hi"
# → helloworld

curl "localhost:8080/?msg=hello"
# → error: expected 'hi' (400 Bad Request)
```

## Project Structure

```
main.go      # Entry point
server.go    # Server initialization and startup
handler.go   # HTTP request handlers
```
