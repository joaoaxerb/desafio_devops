package rotas

import (
    "encoding/json"
    "log"
    "net/http"
    "time"
)

func TextoHandler(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    
    response := Response{
        Message: "Bem vindo ao desafio técnico DevOps Globo!",
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
    
    elapsed := time.Since(start).Milliseconds()
    log.Printf("TextoHandler executou em %d ms", elapsed)
}
