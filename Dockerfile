FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
COPY vendor/ vendor/
COPY internal/ internal/
COPY sql/ sql/
COPY *.go ./
COPY sqlc.yaml ./

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o questlog .

FROM alpine:3.19

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/questlog .

CMD ["./questlog"]
