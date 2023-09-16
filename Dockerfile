FROM golang:1.16 as modules
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM golang:1.16 as builder
WORKDIR /app
COPY . .
RUN go build -o tempDB

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/tempDB .
EXPOSE 8080
CMD ["./tempDB"]
