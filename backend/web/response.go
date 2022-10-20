package web

import "backend/desafio/database"

// ResponseBodyJSONDefault ...
type ResponseBodyJSONDefault struct {
	CodResponse int    `json:"codResponse"`
	Message     string `json:"message"`
}

type ProdutoresResponse struct {
	ResponseBodyJSONDefault
	Producers []database.Producer `json:"producers"`
}
