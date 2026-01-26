import logging
import time
from functools import wraps

logger = logging.getLogger(__name__)

def tempo_execucao(func):
    @wraps(func)
    async def wrapper(*args, **kwargs):
        inicio = time.time()
        resultado = await func(*args, **kwargs)
        fim = time.time()
        tempo_execucao = (fim - inicio) * 1000
        logger.info(f"{func.__name__} executou em {tempo_execucao:.5f} ms")
        return resultado
    return wrapper
