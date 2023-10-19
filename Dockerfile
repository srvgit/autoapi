
FROM golang:1.21.3-alpine3.18 AS builder
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/autoapi
FROM gcr.io/distroless/base
COPY --from=builder /app/autoapi /app/autoapi
CMD ["/app/autoapi"]