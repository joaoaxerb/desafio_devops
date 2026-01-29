from prometheus_client import Counter, Histogram, Gauge, generate_latest, CONTENT_TYPE_LATEST
from fastapi import APIRouter, Response

router = APIRouter()

requisicoes_total = Counter(
    'app_requisicoes_total',
    'Total de requisições',
    ['metodo', 'endpoint', 'status']
)

latencia_requisicoes = Histogram(
    'app_latencia_requisicoes_segundos',
    'Latência das requisições em segundos',
    ['endpoint']
)

cache_hits = Counter('app_cache_hits_total', 'Total de cache hits')
cache_misses = Counter('app_cache_misses_total', 'Total de cache misses')

@router.get("/metricas")
def metricas():
    return Response(generate_latest(), media_type=CONTENT_TYPE_LATEST)
