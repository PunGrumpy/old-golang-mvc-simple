FROM golang:1.21.0-alpine
WORKDIR /app/
COPY . .
RUN go test -race -coverprofile=coverage.out ./... && \
    CGO_ENABLED=1 go build -o simple-mvc ./cmd/main.go

