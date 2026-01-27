package rotas

// Response estrutura de resposta padrão
type Response struct {
    Message string      `json:"message,omitempty"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}
