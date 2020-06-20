# Cache dependancies as builder
FROM golang:1.14 AS builder
WORKDIR /project
COPY . /project
RUN GOOS=linux GOARCH=amd64 go build -a -o  ./app ./cmd/app/main.go

# Final binary
FROM gcr.io/distroless/base
WORKDIR /app
COPY --from=builder /project/app /app/app
COPY --from=builder /project/templates /app/templates
CMD ["/app/app"]

