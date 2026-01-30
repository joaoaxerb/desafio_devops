#  Desafio DevOps

Projeto de desafio técnico DevOps implementando duas aplicações (Python e Go) com cache Redis e stack completa de observabilidade.

## Índice

- [Sobre o Projeto](#sobre-o-projeto)
- [Arquitetura](#arquitetura)
- [Tecnologias](#tecnologias)
- [Como Executar](#como-executar)
- [Aplicações](#aplicações)
- [Observabilidade](#observabilidade)
- [CI/CD](#cicd)

## Sobre o Projeto

Sistema distribuído com duas aplicações independentes que compartilham uma camada de cache Redis, implementando endpoints REST com tempos de cache configuráveis e monitoramento completo através de métricas Prometheus e dashboards Grafana.

### Funcionalidades

- APIs REST em Python (FastAPI) e Go
- Cache distribuído com Redis
- Observabilidade com Prometheus + Grafana
- Health checks e métricas de performance
- Containerização com Docker
- CI/CD com GitHub Actions

##  Arquitetura

<img width="2964" height="2440" alt="image" src="https://github.com/user-attachments/assets/54a73c17-d7e3-4ccc-9a14-c8bf2b3cfab2" />


#  Tecnologias 

## Backend 
- **Python**: FastAPI, fastapi-cache2, Redis, Prometheus Client, Uvicorn 
- **Go**: net/http, go-redis, Prometheus Client Golang
## Infraestrutura 
- **Docker & Docker Compose**: Orquestração de containers 
- **Redis**: Cache in-memory 
- **Prometheus**: Time-series database para métricas 
- **Grafana**: Plataforma de visualização
## DevOps 
- **GitHub Actions**: CI/CD pipeline 
- **Docker Hub/GHCR**: Registry de imagens ---

# Como Executar 
## Pré-requisitos 
- Docker 20.10+ 
- Docker Compose 2.0+ 
## Executar o Projeto Completo 
1. **Clone o repositório** 

```bash 
git clone https://github.com/joaoaxerb/desafio_devops.git
cd desafio_devops
```
2. **Suba todos os serviços**
```bash
docker compose up -d --build
```
3. ***Verifique o status:***
```bash
docker compose ps
```
4. **Acesse as aplicações:**
- App Python: http://localhost:8000

- App Go: http://localhost:8001

- Grafana: http://localhost:3000 (admin/admin)

- Prometheus: http://localhost:9090

6. **Visualize as métricas no Grafana:**
- Acesse http://localhost:3000
- Login: admin / Senha: admin
- Dashboard: "Monitoramento - Desafio DevOps Globo"

# Aplicações 
## Aplicação Python 
FastAPI com cache Redis de **60 segundos**. 
### Endpoints 
- `GET /texto` → Mensagem de boas-vindas (cache 60s) 
- `GET /hora` → Hora atual do servidor (cache 60s) 
- `GET /health` → Health check 
- `GET /metricas` → Métricas Prometheus 
- `GET /docs` → Documentação Swagger **Documentação completa** --- 
##  Aplicação Go 
Servidor HTTP nativo com cache Redis de **10 segundos**. 
### Endpoints 
- `GET /texto` → Mensagem de boas-vindas (cache 10s) 
- `GET /hora` → Hora atual do servidor (cache 10s) 
- `GET /health` → Health check 
- `GET /metricas` → Métricas Prometheus

# Dashboard Grafana 
O dashboard **"Monitoramento - Desafio DevOps Globo"** exibe: 
## Status das Aplicações 
- Health check (Up/Down) 
- Indicadores visuais de disponibilidade ## Métricas do Redis 
- Clientes conectados 
- Comandos processados por segundo ## Performance das Aplicações 
- Consumo de memória (time series) 
- Latência média por aplicação 
- Requisições por endpoint 
## Taxa de Requisições 
- Requests/segundo agregado por job - Detalhamento por endpoint 
## Análise de Cache 
- Cache hits vs misses 
- Efetividade do cache por aplicação 
--- 
# Acessando Métricas 
### Endpoints de métricas 
- **Python**: [http://localhost:8000/metricas](http://localhost:8000/metricas) 
- **Go**: [http://localhost:8001/metricas](http://localhost:8001/metricas) 
- **Redis**: [http://localhost:9121/metrics](http://localhost:9121/metrics) 
- **Prometheus UI**: [http://localhost:9090](http://localhost:9090) 
- **Grafana**: [http://localhost:3000](http://localhost:3000)

## CI/CD Pipeline 
Automatizado via **GitHub Actions** com 3 jobs principais: 
## 1️ Lint 
- Validação de código **Python** com [Ruff](https://github.com/astral-sh/ruff) 
- Linting **Go** com [golangci-lint](https://golangci-lint.run/) 
## 2️ Security 
- **SAST** com [CodeQL](https://codeql.github.com/) 
- Análise de vulnerabilidades 
## 3 Build & Push 
- Build de imagens **Docker** 
- Push para **GitHub Container Registry (GHCR)** - Tagging automático


