package cache

import (
    "context"
    "log"
    "net/http"
    "time"

    "github.com/joaoaxer/desafio_devops/app_go/metricas"
    "github.com/redis/go-redis/v9"
)

var (
    Client *redis.Client
    Ctx    = context.Background()
)

func InitRedis() {
    Client = redis.NewClient(&redis.Options{
        Addr: "redis:6379",
        DB:   0,
    })

    _, err := Client.Ping(Ctx).Result()
    if err != nil {
        log.Fatal("Erro ao conectar no Redis:", err)
    }
    log.Println("Redis inicializado")
}

func CacheMiddleware(next http.HandlerFunc, ttl time.Duration) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        cacheKey := "go-cache:" + r.URL.Path

        // Tenta buscar do cache
        cached, err := Client.Get(Ctx, cacheKey).Result()
        if err == nil {
            // Cache hit
            metricas.CacheHits.Inc()
            w.Header().Set("Content-Type", "application/json")
            w.Write([]byte(cached))

            duration := time.Since(start).Seconds()
            metricas.LatenciaRequisicoes.WithLabelValues(r.URL.Path).Observe(duration)
            metricas.RequisocoesTotal.WithLabelValues(r.Method, r.URL.Path, "200").Inc()
            
            log.Printf("Cache HIT para %s", r.URL.Path)
            return
        }

         // Cache miss
        metricas.CacheMisses.Inc()

        gravador := &gravadorResposta{
            ResponseWriter: w,
            corpo:          []byte{},
            codigoStatus:   http.StatusOK,
        }

        next(gravador, r)

        // Salva no cache
        if gravador.codigoStatus == http.StatusOK {
            Client.Set(Ctx, cacheKey, gravador.corpo, ttl)
            log.Printf("Cache MISS para %s - salvando por %v", r.URL.Path, ttl)
        }

        // Registra métricas
        duration := time.Since(start).Seconds()
        metricas.LatenciaRequisicoes.WithLabelValues(r.URL.Path).Observe(duration)
        metricas.RequisocoesTotal.WithLabelValues(r.Method, r.URL.Path, "200").Inc()
    }
}

type gravadorResposta struct {
    http.ResponseWriter
    corpo         []byte
    codigoStatus  int
}

func (gr *gravadorResposta) Write(b []byte) (int, error) {
    gr.corpo = append(gr.corpo, b...)
    return gr.ResponseWriter.Write(b)
}

func (gr *gravadorResposta) WriteHeader(codigoStatus int) {
    gr.codigoStatus = codigoStatus
    gr.ResponseWriter.WriteHeader(codigoStatus)
}