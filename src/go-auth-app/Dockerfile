# FROM golang:1.16.5 AS development
FROM golang:1.18-alpine3.15
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install github.com/cespare/reflex@latest
EXPOSE 4000
CMD reflex -g '*.go' go run cmd/web/*.go --start-service