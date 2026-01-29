package metricas

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
)

var (
    RequisocoesTotal = promauto.NewCounterVec(
        prometheus.CounterOpts{
            Name: "app_requisicoes_total",
            Help: "Total de requisições",
        },
        []string{"metodo", "endpoint", "status"},
    )

    LatenciaRequisicoes = promauto.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "app_latencia_requisicoes_segundos",
            Help:    "Latência das requisições em segundos",
            Buckets: prometheus.DefBuckets,
        },
        []string{"endpoint"},
    )

    CacheHits = promauto.NewCounter(
        prometheus.CounterOpts{
            Name: "app_cache_hits_total",
            Help: "Total de cache hits",
        },
    )

    CacheMisses = promauto.NewCounter(
        prometheus.CounterOpts{
            Name: "app_cache_misses_total",
            Help: "Total de cache misses",
        },
    )
)
