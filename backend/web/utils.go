package web

import (
	"backend/desafio/database"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func handleTransactionData(data string) ([]*database.Producer, error) {
	var err error
	lines := strings.Split(data, "\n")
	dictProd := map[string]*database.Producer{}
	producers := make([]*database.Producer, 0)
	for index, line := range lines {

		producer := &database.Producer{}
		if len(line) == 0 {
			break
		}

		name := strings.TrimSpace(line[66:])
		if _, ok := dictProd[name]; !ok {
			producer.Name = name
			dictProd[name] = producer
			producer.Transactions = make([]database.Transaction, 0)
			producers = append(producers, producer)
		} else {
			producer = dictProd[name]
		}
		transactions := database.Transaction{}
		transactions.TypeID, err = strconv.Atoi(line[0:1])
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Invalid type value format on line %d", index+1))
		}

		transactions.Date, err = time.Parse(time.RFC3339, line[1:26])
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Invalid date value format in line %d", index+1))
		}

		transactions.Product = strings.TrimSpace(line[26:56])
		valorFull := line[56:66]
		valorFull = valorFull[0:len(valorFull)-2] + "." + valorFull[len(valorFull)-2:]
		transactions.Value, err = strconv.ParseFloat(valorFull, 64)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Invalid transaction amount format on line %d", index+1))
		}

		producer.Transactions = append(producer.Transactions, transactions)

	}
	return producers, nil
}
