package rotas

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"
)

func HoraHandler(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    
    horaAtual := time.Now().Format("02/01/2006 15:04:05")
    
    response := Response{
        Message: fmt.Sprintf("Hora do servidor: %s", horaAtual),
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
    
    elapsed := time.Since(start)
    log.Printf("HoraHandler executou em %.5f ms", elapsed.Seconds()*1000)
}
