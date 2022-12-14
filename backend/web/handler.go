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

func (h *Handler) CreateTransactions(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	var buf bytes.Buffer
	file, _, err := r.FormFile("file")
	if err != nil {
		h.CoreRespondErro(w, r, "", errors.New("Invalid format file"), "Error on create transaction", http.StatusBadRequest)
		return
	}

	defer file.Close()
	io.Copy(&buf, file)
	contents := buf.String()

	produtores, err := handleTransactionData(contents)
	if err != nil {
		h.CoreRespondErro(w, r, "", err, "Error on unpack file", http.StatusBadRequest)
		return
	}

	err = (*h.Database).CreateTransaction(produtores)
	if err != nil {
		h.CoreRespondErro(w, r, "", errors.New("It's is not possible to save the data internal error"), "Error on create transaction", http.StatusInternalServerError)
		return
	}

	bit, _ := json.Marshal(produtores)
	fmt.Println(string(bit))
	h.CoreRespondSucess(w, r, http.StatusCreated, ResponseBodyJSONDefault{CodResponse: http.StatusCreated, Message: "Transactions Created"})
	return

}

func (h *Handler) ListarTransacoes(w http.ResponseWriter, r *http.Request) {

	producers, err := (*h.Database).ListTransaction()
	if err != nil {
		fmt.Println(err)
		h.CoreRespondErro(w, r, "", errors.New("There is no transactions"), "Error on retrieve transactions", http.StatusInternalServerError)
		return
	}
	producers = sumProducersTransactions(producers)

	bit, _ := json.Marshal(producers)
	fmt.Println(string(bit))
	response := ProdutoresResponse{
		Producers: producers,
		ResponseBodyJSONDefault: ResponseBodyJSONDefault{
			CodResponse: http.StatusOK,
			Message:     "Transações retornadas com sucesso",
		},
	}
	response.CodResponse = http.StatusOK
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
