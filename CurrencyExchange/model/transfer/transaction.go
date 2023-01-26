package transfer

import (
	"root/model/account"
	"root/model/currency"
)

func Transaction(accounts *[]account.Account, transfers *[]Transfer) {
	for index := range *accounts {
		for _, transaction := range (*transfers)[index].Transaction {
			currency.FromTo(&(*accounts)[index].Currency, transaction.From, transaction.To, transaction.Deposit)
		}
	}
}
