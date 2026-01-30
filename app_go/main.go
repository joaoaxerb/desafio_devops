package main

import (
    "log"
    "net/http"
	"time"

    "github.com/joaoaxer/desafio_devops/app_go/cache"
    "github.com/joaoaxer/desafio_devops/app_go/rotas"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	cache.InitRedis()

	// Rotas sem cache
	http.HandleFunc("/health", rotas.HealthHandler)

	// Rotas com cache
    http.HandleFunc("/texto", cache.CacheMiddleware(rotas.TextoHandler, 10*time.Second))
    http.HandleFunc("/hora", cache.CacheMiddleware(rotas.HoraHandler, 10*time.Second))

	// Mtricas
	http.Handle("/metricas", promhttp.Handler())

	port := "8001"
	log.Printf("Servidor Go iniciado na porta %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}