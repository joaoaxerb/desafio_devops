# Aplicação Go - HTTP Server

API REST desenvolvida com servidor HTTP nativo do Go, implementando cache Redis com TTL de 10 segundos e exportação de métricas Prometheus.

## Tecnologias

- **Go 1.23**
- **net/http**: Servidor HTTP padrão da biblioteca
- **go-redis/v9**: Client Redis oficial
- **Prometheus Client Golang**: Exportação de métricas

## Como Executar

### Com Docker (Recomendado)

```bash
# No diretório raiz do projeto
docker compose up app-go -d
```

### Desenvolvimento Local

1. **Instalar dependências:**
   ```bash
   cd app_go
   go mod download
   ```

2. **Configurar Redis:**
   ```bash
   docker run -d -p 6379:6379 redis:7-alpine
   ```

3. **Executar:**
   ```bash
   go run main.go
   ```

4. **Build:**
   ```bash
   go build -o app main.go
   ./app
   ```

## Endpoints

### GET /texto

Retorna mensagem de boas-vindas com cache de 10 segundos.

**Request:**
```bash
curl http://localhost:8001/texto
```

**Response:**
```json
{
  "mensagem": "Bem vindo ao desafio técnico DevOps Globo!",
  "timestamp": "2026-01-30T15:30:45Z"
}
```

### GET /hora

Retorna hora atual do servidor com cache de 10 segundos.

**Request:**
```bash
curl http://localhost:8001/hora
```

**Response:**
```json
{
  "hora": "30/01/2026 15:30:45",
  "timestamp": "2026-01-30T15:30:45Z"
}
```

### GET /health

Health check da aplicação.

**Request:**
```bash
curl http://localhost:8001/health
```

**Response:**
```json
{
  "status": "healthy"
}
```

### GET /metricas

Métricas no formato Prometheus.

**Request:**
```bash
curl http://localhost:8001/metricas
```

**Response:**
```
# HELP app_requisicoes_total Total de requisicoes
# TYPE app_requisicoes_total counter
app_requisicoes_total{endpoint="/texto",metodo="GET",status="200"} 42

# HELP app_latencia_requisicoes_segundos Latencia das requisicoes
# TYPE app_latencia_requisicoes_segundos histogram
app_latencia_requisicoes_segundos_bucket{endpoint="/texto",le="0.005"} 40

# HELP app_cache_hits_total Total de cache hits
# TYPE app_cache_hits_total counter
app_cache_hits_total{endpoint="/texto"} 38

# HELP app_cache_misses_total Total de cache misses
# TYPE app_cache_misses_total counter
app_cache_misses_total{endpoint="/texto"} 4
```

## Métricas

| Métrica | Tipo | Descrição | Labels |
|---------|------|-----------|--------|
| `app_requisicoes_total` | Counter | Total de requisições HTTP | `metodo`, `endpoint`, `status` |
| `app_latencia_requisicoes_segundos` | Histogram | Latência das requisições | `endpoint` |
| `app_cache_hits_total` | Counter | Cache hits | `endpoint` |
| `app_cache_misses_total` | Counter | Cache misses | `endpoint` |

As métricas são coletadas automaticamente pelo `CacheMiddleware` que intercepta todas as requisições e registra:
- Hits e misses do cache
- Latência de resposta
- Status HTTP
- Endpoint acessado

## Configuração
### Cache

Todos os endpoints possuem TTL de **10 segundos**:

| Endpoint | TTL | 
|----------|-----|
| `/texto` | 10s |
| `/hora` | 10s |

**Configuração no código:**

```go
http.HandleFunc("/texto", cache.CacheMiddleware(rotas.TextoHandler, 10*time.Second))
http.HandleFunc("/hora", cache.CacheMiddleware(rotas.HoraHandler, 10*time.Second))
```

## Estrutura

```
app_go/
├── main.go              # Entry point
├── rotas/
│   ├── texto.go        # Handler /texto
│   ├── hora.go         # Handler /hora
│   └── health.go       # Handler /health
├── cache/
│   └── redis.go        # Middleware de cache + client Redis
├── metricas/
│   └── metricas.go     # Definição das métricas
├── Dockerfile
├── go.mod
└── README.md
```

## Docker

### Build

```bash
docker build -t app-go:latest ./app_go
```

### Run

```bash
docker run -d \
  --name app-go \
  -p 8001:8001 \
  -e REDIS_URL=redis://redis:6379 \
  app-go:latest
```

### Multi-stage Build

O Dockerfile usa multi-stage build para otimizar o tamanho:

```dockerfile
# Stage 1: Build
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

# Stage 2: Run
FROM golang:1.23-alpine
WORKDIR /app
COPY --from=builder /app/main .
USER appuser
CMD ["./main"]
```

## Testando o Cache

Observe o comportamento do cache com TTL de 10 segundos:

```bash
# Primeira requisição (MISS)
time curl http://localhost:8001/hora
# Latência: ~40-80ms

# Segunda requisição (HIT)
time curl http://localhost:8001/hora
# Latência: ~3-8ms (muito mais rápida!)

# Aguarde 10 segundos para expirar
sleep 10

# Após expiração (MISS novamente)
time curl http://localhost:8001/hora
# Latência: ~40-80ms
```

**Verificar métricas:**
```bash
curl http://localhost:8001/metricas | grep cache

# Output:
# app_cache_hits_total{endpoint="/hora"} 5
# app_cache_misses_total{endpoint="/hora"} 2
```

## Logs

```bash
# Ver logs em tempo real
docker compose logs -f app-go

# Últimas 50 linhas
docker compose logs --tail=50 app-go
```

**Exemplo:**
```
2026/01/30 15:30:45 Servidor Go iniciado na porta 8001
2026/01/30 15:30:50 Redis conectado com sucesso
2026/01/30 15:31:00 Cache miss para /texto
2026/01/30 15:31:05 Cache hit para /texto
```

## Debugging

### Verificar Conexão Redis

```bash
# Conectar ao Redis
docker compose exec redis redis-cli

# Listar chaves da aplicação Go
KEYS app-go:*

# Output:
# 1) "app-go:/texto"
# 2) "app-go:/hora"

# Ver conteúdo
GET app-go:/texto

# Ver TTL restante
TTL app-go:/texto
# Output: 8 (segundos restantes)

# Limpar cache
DEL app-go:/texto app-go:/hora
```

### Testar Endpoint de Saúde

```bash
# Health check
curl http://localhost:8001/health

# Verificar processo
docker compose exec app-go ps aux
```

## Dependências

```go
module github.com/joaoaxer/desafio_devops/app_go

go 1.23

require (
    github.com/redis/go-redis/v9 v9.7.0
    github.com/prometheus/client_golang v1.20.5
)
```

## Atualizar Dependências

```bash
# Atualizar todas as dependências
go get -u ./...
go mod tidy
```

## Troubleshooting

### Erro: "Cannot connect to Redis"

```bash
# Verificar se Redis está rodando
docker compose ps redis

# Ver logs
docker compose logs redis

# Testar conexão
docker compose exec redis redis-cli ping
```

### Porta já em uso

```bash
# Verificar processo na porta 8001
lsof -i :8001

# Matar processo
kill -9 <PID>
```

### Métricas não aparecem

```bash
# Verificar endpoint
curl http://localhost:8001/metricas

# Verificar Prometheus targets
# http://localhost:9090/targets
```

## Boas Práticas

- **Async Redis**: Operações Redis não bloqueantes
- **Context Timeout**: Timeout configurado para operações
- **Error Handling**: Tratamento robusto de erros
- **Structured Logging**: Logs formatados
- **Non-root User**: Container roda com usuário não-root
- **Multi-stage Build**: Imagem otimizada

---

**[Voltar para documentação principal](../README.md)**
