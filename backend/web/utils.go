package web

import (
	"backend/desafio/database"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func handleTransactionData(data string) ([]*database.Produtor, error) {
	var err error
	lines := strings.Split(data, "\n")
	dictProd := map[string]*database.Produtor{}
	produtores := make([]*database.Produtor, 0)
	for index, line := range lines {
		fmt.Printf("Linha - %d | %s | %d\n", index+1, line, len(line))
		produtor := &database.Produtor{}
		if len(line) == 0 {
			break
		}
		// if len(line) != 86 {
		// 	return nil, errors.New(fmt.Sprintf("Formato do arquivo inválido na linha %d com tamanho de caracteres diferente de 86", index+1))
		// }
		nome := strings.TrimSpace(line[66:])
		if _, ok := dictProd[nome]; !ok {
			produtor.Nome = nome
			dictProd[nome] = produtor
			produtor.Transacoes = make([]database.Transacao, 0)
			produtores = append(produtores, produtor)
		} else {
			produtor = dictProd[nome]
		}
		transacao := database.Transacao{}
		transacao.TipoID, err = strconv.Atoi(line[0:1])
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Formato do valor do tipo inválido na linha %d", index+1))
		}

		transacao.Data, err = time.Parse(time.RFC3339, line[1:26])
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Formato do valor da data inválida na linha %d", index+1))
		}

		transacao.Produto = strings.TrimSpace(line[26:56])
		valorFull := line[56:66]
		valorFull = valorFull[0:len(valorFull)-2] + "." + valorFull[len(valorFull)-2:]
		transacao.Valor, err = strconv.ParseFloat(valorFull, 64)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Formato do valor da transação inválido na linha %d", index+1))
		}

		produtor.Transacoes = append(produtor.Transacoes, transacao)

	}
	return produtores, nil
}
