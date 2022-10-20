package web

import (
	"backend/desafio/database"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"testing"
)

func initDataBase() (database.IDataBase, error) {
	DBOpt := &database.OptionsDBClient{
		URL:    "postgresql://root:luke@localhost:5434/app",
		Driver: "postgres",
	}
	db, err := DBOpt.ConfigDatabase()
	return *db, err

}
func initHandler() *Handler {
	handler := &Handler{}

	DB, errDB := initDataBase()
	if errDB != nil {
		fmt.Println(errDB)
	}
	handler.Database = &DB

	return handler
}

func TestCadastrarTransacao(t *testing.T) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, _ := writer.CreateFormFile("file", "../../sales.txt")
	file, _ := os.Open("../../sales.txt")
	io.Copy(fw, file)
	writer.Close()
	r := httptest.NewRequest("POST", "/transactions", bytes.NewReader(body.Bytes()))
	r.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	hWeb := initHandler()
	RouterTest(hWeb).ServeHTTP(w, r)
	_, errBody := ioutil.ReadAll(w.Body)
	if errBody != nil {
		t.Error("expect nil, got - ", errBody)
	}
}

func TestListTransacoes(t *testing.T) {
	hWeb := initHandler()
	r := httptest.NewRequest("GET", "/transactions", nil)
	w := httptest.NewRecorder()
	RouterTest(hWeb).ServeHTTP(w, r)
	data, errBody := ioutil.ReadAll(w.Body)
	fmt.Println(string(data))
	if errBody == nil {
		t.Error("expect nil, got - ", errBody)
	}
}
