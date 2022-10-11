package web

import (
	"backend/desafio/core"
	"backend/desafio/database"
	"bytes"
	"encoding/json"
	"fmt"

	"errors"
	"io"

	"net/http"
	"strconv"
)

type Handler struct {
	Database *database.IDataBase
}

func (h *Handler) CadastrarTransacoes(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	var buf bytes.Buffer
	file, _, err := r.FormFile("file")
	if err != nil {
		h.CoreRespondErro(w, r, "", errors.New("Arquivo com formato inválido"), "ErrCadPlaneta", http.StatusBadRequest)
		return
	}

	defer file.Close()
	io.Copy(&buf, file)
	contents := buf.String()
	//fmt.Println("DEPOIS -------- ", contents)
	produtores, err := handleTransactionData(contents)
	if err != nil {
		h.CoreRespondErro(w, r, "", err, "Erro ao descompactar", http.StatusBadRequest)
		return
	}
	//db := *h.Database

	err = (*h.Database).CreateTransacao(produtores)
	if err != nil {
		h.CoreRespondErro(w, r, "", errors.New("Arquivo com formato inválido"), "ErrCadPlaneta", http.StatusBadRequest)
		return
	}

	bit, _ := json.Marshal(produtores)
	fmt.Println(string(bit))
	h.CoreRespondSucess(w, r, http.StatusCreated, ResponseBodyJSONDefault{CodResposta: http.StatusCreated, Mensagem: "Transações Cadastradas"})
	return

}

func (h *Handler) ListarTransacoes(w http.ResponseWriter, r *http.Request) {

	produtores, err := (*h.Database).ListTransacao()
	if err != nil {
		h.CoreRespondErro(w, r, "", err, "Erro ao listar transacoes", http.StatusInternalServerError)
		return
	}
	response := ProdutoresResponse{
		Produtores: produtores,
		ResponseBodyJSONDefault: ResponseBodyJSONDefault{
			CodResposta: http.StatusOK,
			Mensagem:    "Transações retornadas com sucesso",
		},
	}
	response.CodResposta = http.StatusOK
	h.CoreRespondSucess(w, r, http.StatusOK, response)
	return

}

func (h *Handler) CoreRespondErro(w http.ResponseWriter, r *http.Request, idOperation string, erro error, message string, codeError int) {
	core.Respond(w, r, codeError, core.ErrDetail{
		Resource: message, Code: strconv.Itoa(codeError), Message: erro.Error(), IDOperation: idOperation,
	})
}

func (h *Handler) CoreRespondSucess(w http.ResponseWriter, r *http.Request, code int, object interface{}) {
	core.Respond(w, r, code, object)
}
