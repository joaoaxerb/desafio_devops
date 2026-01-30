# Aplicação Python - FastAPI

API REST desenvolvida com FastAPI, implementando cache Redis com TTL de 60 segundos e coleta de métricas Prometheus.

## Tecnologias

- **Python 3.11**
- **FastAPI**: Framework web assíncrono de alta performance
- **fastapi-cache2**: Biblioteca de cache com suporte Redis
- **Redis**: Client assíncrono (aioredis)
- **Prometheus Client**: Exportação de métricas
- **Uvicorn**: Servidor ASGI

## Como Executar

### Com Docker (Recomendado)

```bash
# No diretório raiz do projeto
docker compose up app-python -d
```

### Desenvolvimento Local

1. **Instalar dependências:**
   ```bash
   cd app_python
   
   # Com uv (recomendado)
   uv sync
   
   # Ou com pip
   pip install -r requirements.txt
   ```

2. **Configurar Redis:**
   ```bash
   # Certifique-se que o Redis está rodando
   docker run -d -p 6379:6379 redis:7-alpine
   ```

3. **Executar aplicação:**
   ```bash
   # Com uv
   uv run uvicorn main:app --app-dir src --reload --host 0.0.0.0 --port 8000
   
   # Ou diretamente
   uvicorn main:app --app-dir src --reload --host 0.0.0.0 --port 8000
   ```

## Endpoints

### API Principal

#### GET /texto

Retorna mensagem de boas-vindas com cache de **60 segundos**.

**Request:**
```bash
curl http://localhost:8000/texto
```

**Response:**
```json
{
  "status": "sucesso",
  "dados": "Bem vindo ao desafio técnico DevOps Globo!",
  "timestamp": "2026-01-30T15:30:45"
}
```

#### GET /hora

Retorna hora atual do servidor com cache de **60 segundos**.

**Request:**
```bash
curl http://localhost:8000/hora
```

**Response:**
```json
{
  "status": "sucesso",
  "dados": "Hora do servidor: 30/01/2026 15:30:45",
  "timestamp": "2026-01-30T15:30:45"
}
```

#### GET /health

Health check da aplicação.

**Request:**
```bash
curl http://localhost:8000/health
```

**Response:**
```json
{
  "status": "healthy"
}
```

### Documentação e Métricas

#### GET /docs

Documentação Swagger UI interativa.

**Acesso:** http://localhost:8000/docs

#### GET /redoc

Documentação ReDoc alternativa.

**Acesso:** http://localhost:8000/redoc

#### GET /metricas

Métricas no formato Prometheus.

**Request:**
```bash
curl http://localhost:8000/metricas
```

**Response:**
```
# HELP app_requisicoes_total Total de requisicoes
# TYPE app_requisicoes_total counter
app_requisicoes_total{endpoint="/texto",metodo="GET",status="200"} 42.0
app_requisicoes_total{endpoint="/hora",metodo="GET",status="200"} 38.0

# HELP app_latencia_requisicoes_segundos Latencia das requisicoes
# TYPE app_latencia_requisicoes_segundos histogram
app_latencia_requisicoes_segundos_bucket{endpoint="/texto",le="0.005"} 40.0
app_latencia_requisicoes_segundos_bucket{endpoint="/texto",le="0.01"} 41.0
app_latencia_requisicoes_segundos_sum{endpoint="/texto"} 0.234
app_latencia_requisicoes_segundos_count{endpoint="/texto"} 42.0

# HELP app_cache_hits_total Total de cache hits
# TYPE app_cache_hits_total counter
app_cache_hits_total{endpoint="/texto"} 38.0

# HELP app_cache_misses_total Total de cache misses
# TYPE app_cache_misses_total counter
app_cache_misses_total{endpoint="/texto"} 4.0
```

## Métricas Implementadas

| Métrica | Tipo | Descrição | Labels |
|---------|------|-----------|--------|
| `app_requisicoes_total` | Counter | Total de requisições HTTP | `metodo`, `endpoint`, `status` |
| `app_latencia_requisicoes_segundos` | Histogram | Tempo de resposta das requisições | `endpoint` |
| `app_cache_hits_total` | Counter | Número de cache hits | `endpoint` |
| `app_cache_misses_total` | Counter | Número de cache misses | `endpoint` |

### Coleta Automática

As métricas são coletadas automaticamente pelo middleware `MetricasMiddleware`, que intercepta todas as requisições HTTP e registra:

- Método HTTP (GET, POST, etc)
- Endpoint acessado
- Status code da resposta
- Tempo de execução (latência em segundos)

## Configuração

### Cache

O cache é configurado com TTL de **60 segundos** para ambos os endpoints:

| Endpoint | TTL | Comportamento |
|----------|-----|---------------|
| `/texto` | 60s | Conteúdo estático - cache efetivo |
| `/hora` | 60s | Hora atualizada a cada minuto |

**Implementação:**

```python
@router.get("/texto")
@cache(expire=60)  # Cache de 60 segundos
async def texto_fixo(request: Request):
    return resposta_sucesso("Bem vindo ao desafio técnico DevOps Globo!")
```

## Estrutura do Código

```
app_python/
├── src/
│   ├── main.py                    # Entry point da aplicação
│   ├── rotas/
│   │   ├── __init__.py           # Router principal
│   │   ├── texto_fixo.py         # Endpoint /texto
│   │   ├── hora_servidor.py      # Endpoint /hora
│   │   └── metricas.py           # Endpoint /metricas + definições
│   ├── middleware/
│   │   └── metricas_middleware.py # Coleta automática de métricas
│   └── utils/
│       ├── decorators.py         # Decoradores auxiliares
│       └── respostas.py          # Formatação de respostas
├── Dockerfile
├── pyproject.toml
├── requirements.txt
└── README.md
```

### Componentes Principais

#### main.py

Entry point da aplicação que:
- Inicializa o FastAPI
- Configura o cache Redis
- Registra o middleware de métricas
- Inclui os routers

#### middleware/metricas_middleware.py

Middleware que coleta métricas automaticamente em todas as requisições:
- Incrementa contador de requisições por endpoint/método/status
- Registra latência em histograma
- Não interfere no fluxo da aplicação

#### rotas/metricas.py

Define as métricas Prometheus:
- `requisicoes_total`: Counter com labels
- `latencia_requisicoes`: Histogram
- `cache_hits_total` / `cache_misses_total`: Counters

## Docker

### Build

```bash
docker build -t app-python:latest ./app_python
```

### Run

```bash
docker run -d \
  --name app-python \
  -p 8000:8000 \
  -e REDIS_URL=redis://redis:6379 \
  app-python:latest
```

### Segurança

O Dockerfile implementa best practices:
- Usuário não-root (`appuser`)
- Minimal base image (python:3.11-slim)
- Dependencies cacheadas em layer separado
- Ownership adequado dos arquivos

## Testando o Cache

Observe o comportamento do cache fazendo requisições consecutivas:

```bash
# Primeira requisição (MISS) - consulta o backend
time curl http://localhost:8000/texto
# Latência: ~50-100ms

# Segunda requisição (HIT) - retorna do cache
time curl http://localhost:8000/texto
# Latência: ~5-10ms (muito mais rápida!)

# Requisições subsequentes dentro de 60s (HIT)
for i in {1..5}; do
  curl -s http://localhost:8000/texto | jq '.timestamp'
done
# Timestamp permanece o mesmo (resposta cacheada)

# Aguarde 60 segundos para o cache expirar
sleep 60

# Após expiração, volta a ser MISS
time curl http://localhost:8000/texto
# Latência: ~50-100ms novamente
```

### Visualizar Métricas de Cache

```bash
# Verificar cache hits e misses
curl http://localhost:8000/metricas | grep cache

# Output:
# app_cache_hits_total{endpoint="/texto"} 38.0
# app_cache_misses_total{endpoint="/texto"} 4.0
```

**Acompanhar no Grafana:**
- Acesse: http://localhost:3000
- Dashboard: "Monitoramento - Desafio DevOps Globo"
- Painel: "Cache Hits vs Misses"

## Logs

A aplicação usa logging estruturado com saída para stdout:

```bash
# Ver logs em tempo real
docker compose logs -f app-python

# Últimas 100 linhas
docker compose logs --tail=100 app-python
```

**Formato dos logs:**
```
2026-01-30 15:30:45 - uvicorn.access - INFO - 172.20.0.1:45678 - "GET /texto HTTP/1.1" 200
2026-01-30 15:30:45 - root - INFO - Cache Redis inicializado
2026-01-30 15:30:50 - uvicorn.error - INFO - Application startup complete.
```

## Debugging

### Verificar Conexão Redis

Adicione endpoint temporário para testar conexão:

```python
from redis import asyncio as aioredis

@app.get("/redis-test")
async def test_redis():
    redis = aioredis.from_url("redis://redis:6379")
    await redis.set("test", "ok")
    value = await redis.get("test")
    return {"redis_status": value.decode() if value else "error"}
```

### Inspecionar Cache no Redis

```bash
# Conectar ao container Redis
docker compose exec redis redis-cli

# Listar todas as chaves
KEYS *

# Ver chaves do cache da aplicação Python
KEYS fastapi-cache:*

# Output:
# 1) "fastapi-cache:/texto"
# 2) "fastapi-cache:/hora"

# Ver conteúdo de uma chave
GET fastapi-cache:/texto

# Ver TTL restante
TTL fastapi-cache:/texto
# Output: 45 (segundos restantes)

# Limpar cache manualmente
DEL fastapi-cache:/texto
DEL fastapi-cache:/hora

# Ou limpar tudo
FLUSHALL
```

### Verificar Saúde da Aplicação

```bash
# Health check
curl http://localhost:8000/health

# Verificar métricas do processo
curl http://localhost:8000/metricas | grep process_

# Output inclui:
# process_resident_memory_bytes
# process_cpu_seconds_total
# process_open_fds
```

## Dependências

### pyproject.toml

```toml
[project]
name = "app-python"
version = "1.0.0"
requires-python = ">=3.11"
dependencies = [
    "fastapi>=0.115.6",
    "uvicorn>=0.34.0",
    "fastapi-cache2[redis]>=0.2.1",
    "redis>=4.2.0,<5.0.0",
    "prometheus-client>=0.21.1",
]
```

### Atualizar Dependências

```bash
# Com uv
uv sync --upgrade

# Com pip
pip install --upgrade -r requirements.txt
```

## Troubleshooting

### Erro: "Cannot connect to Redis"

```bash
# Verificar se Redis está rodando
docker compose ps redis

# Ver logs do Redis
docker compose logs redis

# Testar conexão manual
docker compose exec redis redis-cli ping
# Output esperado: PONG
```

### Erro: "ModuleNotFoundError: No module named 'rotas'"

```bash
# Executar com --app-dir correto
uvicorn main:app --app-dir src --reload

# Ou ajustar PYTHONPATH
export PYTHONPATH="${PYTHONPATH}:./src"
python -m uvicorn main:app --reload
```

### Métricas não aparecem no Prometheus

```bash
# Verificar se endpoint está acessível
curl http://localhost:8000/metricas

# Verificar configuração do Prometheus
docker compose exec prometheus cat /etc/prometheus/prometheus.yml

# Ver targets no Prometheus
# Acesse: http://localhost:9090/targets
```

## Boas Práticas Implementadas

- **Async/Await**: Toda a stack é assíncrona para melhor performance
- **Type Hints**: Código com type annotations para melhor manutenção
- **Middleware**: Coleta de métricas não invasiva
- **Cache Decorator**: Sintaxe declarativa para cache
- **Structured Logging**: Logs formatados e informativos
- **Health Checks**: Endpoint dedicado para monitoramento
- **Documentation**: Swagger UI automático
- **Security**: Usuário não-root no Docker

---

**[Voltar para documentação principal](../README.md)**
