from typing import Any, Dict, Optional, Union
from fastapi import HTTPException

def resposta_sucesso(data: Any = None) -> Any:
    return data


def resposta_erro(mensagem: str, status_code: int = 500):
    raise HTTPException(status_code=status_code, detail=mensagem)