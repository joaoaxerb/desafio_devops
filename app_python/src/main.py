from typing import Union
from rotas import api_router
from fastapi import FastAPI
import logging
from fastapi_cache import FastAPICache
from fastapi_cache.backends.redis import RedisBackend
from redis import asyncio as aioredis
from middleware.metricas_middleware import MetricasMiddleware

logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s',
    handlers=[
        logging.StreamHandler()
    ]
)

app = FastAPI(
    title="Desafio DevOps API",
    description="API para o desafio técnico DevOps",
    version="1.0.0"
)

app.add_middleware(MetricasMiddleware)

@app.on_event("startup")
async def startup():
    redis = aioredis.from_url("redis://redis:6379" )
    FastAPICache.init(RedisBackend(redis), prefix="fastapi-cache")
    logging.info("Cache Redis inicializado")

app.include_router(api_router)

@app.get("/health")
def health_check():
    return {"status": "healthy"}

