from fastapi import APIRouter
from rotas import hora_servidor
from rotas import texto_fixo
from rotas import metricas

api_router = APIRouter()
api_router.include_router(hora_servidor.router, tags=["hora_servidor"])
api_router.include_router(texto_fixo.router, tags=["texto_fixo"])
api_router.include_router(metricas.router, tags=["metricas"])