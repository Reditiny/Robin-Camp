FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/assignment

FROM alpine:latest

WORKDIR /app

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

COPY --from=builder /app/assignment /app/assignment
COPY --from=builder /app/.env /app/.env

EXPOSE 8080
CMD ["/app/assignment"]