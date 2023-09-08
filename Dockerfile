FROM golang:1.21.0-alpine
WORKDIR /app/
COPY . .
RUN CGO_ENABLED=1 go test -race -coverprofile=coverage.out ./... && \
    go build -o simple-mvc ./cmd/main.go
