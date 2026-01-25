from typing import Union
from rotas import api_router
from fastapi import FastAPI
import logging

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

app.include_router(api_router)

@app.get("/health")
def health_check():
    return {"status": "healthy"}

