FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache ca-certificates git

COPY go.mod go.sum* ./

RUN go mod download || go mod download

COPY . .

RUN go build -o main .

# Reutilizar a mesma imagem base (já está no cache)
FROM golang:1.23-alpine

RUN addgroup -S appuser && adduser -S appuser -G appuser

WORKDIR /app

COPY --from=builder /app/main .

RUN chown -R appuser:appuser /app

USER appuser

EXPOSE 8001

CMD ["./main"]