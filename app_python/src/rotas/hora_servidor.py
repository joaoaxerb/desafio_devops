from fastapi import APIRouter
from datetime import datetime
from utils.decorators import tempo_execucao
from utils.respostas import resposta_sucesso, resposta_erro
from fastapi_cache.decorator import cache

router = APIRouter()

@router.get("/hora")
@cache(expire=60)
@tempo_execucao
async def hora_servidor():
    try:
        hora_atual = datetime.now().strftime("%d/%m/%Y %H:%M:%S")
        return resposta_sucesso(f"Hora do servidor: {hora_atual}")
    except Exception as e:
        return resposta_erro(str(e))