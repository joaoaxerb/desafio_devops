#!/bin/bash

# Script para gerar tráfego nas aplicações Python e Go

TOTAL_REQUESTS="${1:-50}"

echo "Gerando tráfego para popular métricas..."
echo "Total de requisições por endpoint: $TOTAL_REQUESTS"
echo ""

# App Python
echo "Requisições para App Python:"
for i in $(seq 1 $TOTAL_REQUESTS); do
    curl -s http://localhost:8000/texto > /dev/null
    echo -n "."
done
echo " /texto OK"

for i in $(seq 1 $TOTAL_REQUESTS); do
    curl -s http://localhost:8000/hora > /dev/null
    echo -n "."
done
echo " /hora OK"

# App Go
echo ""
echo "Requisições para App Go:"
for i in $(seq 1 $TOTAL_REQUESTS); do
    curl -s http://localhost:8001/texto > /dev/null
    echo -n "."
done
echo " /texto OK"

for i in $(seq 1 $TOTAL_REQUESTS); do
    curl -s http://localhost:8001/hora > /dev/null
    echo -n "."
done
echo " /hora OK"

echo ""
echo "Tráfego gerado com sucesso!"
echo "Total de requisições: $((TOTAL_REQUESTS * 4))"
echo ""
echo "Visualize as métricas em:"
echo "  Grafana: http://localhost:3000"
echo "  Prometheus: http://localhost:9090"