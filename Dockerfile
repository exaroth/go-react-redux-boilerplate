# Cache dependancies as builder
FROM golang:1.14 AS builder
WORKDIR /project
COPY . /project
RUN go get -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

# Final binary
FROM gcr.io/distroless/base
WORKDIR /app
COPY --from=builder /project/app /app/app
CMD ["/app/app"]

