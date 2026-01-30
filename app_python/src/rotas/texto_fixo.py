from fastapi import APIRouter, Request
from utils.decorators import tempo_execucao
from utils.respostas import resposta_sucesso, resposta_erro
from fastapi_cache.decorator import cache

router = APIRouter()

@router.get("/texto")
@cache(expire=60)
@tempo_execucao
async def texto_fixo(request: Request):
    try:
        return resposta_sucesso("Bem vindo ao desafio técnico DevOps Globo!")
    except Exception as e:
        return resposta_erro(str(e))    