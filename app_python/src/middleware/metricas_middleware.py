import time
from starlette.middleware.base import BaseHTTPMiddleware
from rotas.metricas import requisicoes_total, latencia_requisicoes

class MetricasMiddleware(BaseHTTPMiddleware):
    async def dispatch(self, request, call_next):
        # Ignora o próprio endpoint de métricas
        if request.url.path == "/metricas":
            return await call_next(request)
        
        start_time = time.time()
        
        # Processa a requisição
        response = await call_next(request)
        
        # Calcula latência
        duration = time.time() - start_time
        
        # Incrementa métricas
        requisicoes_total.labels(
            metodo=request.method,
            endpoint=request.url.path,
            status=response.status_code
        ).inc()
        
        latencia_requisicoes.labels(
            endpoint=request.url.path
        ).observe(duration)
        
        return response
