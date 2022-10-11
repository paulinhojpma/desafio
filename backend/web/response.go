package web

import "backend/desafio/database"

// ResponseBodyJSONDefault ...
type ResponseBodyJSONDefault struct {
	CodResposta int    `json:"codResposta"`
	Mensagem    string `json:"mensagem"`
}

type ProdutoresResponse struct {
	ResponseBodyJSONDefault
	Produtores []database.Produtor
}
