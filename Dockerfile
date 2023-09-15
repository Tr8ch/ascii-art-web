FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o ascii ./cmd/web/main.go
FROM alpine
WORKDIR /app
COPY --from=builder /app/ascii ./
COPY /internal /app/internal
COPY /static /app/static
CMD ["./ascii"]