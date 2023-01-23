package transfer

import (
	"root/model/person"
	"root/model/transaction"
)

type Transfer struct {
	User        person.Person             `json:"User"`
	Transaction []transaction.Transaction `json:"Transaction"`
}

func Empty() *[]Transfer {
	return &[]Transfer{}
}

func New(User person.Person, Transaction []transaction.Transaction) *Transfer {
	return &Transfer{
		User, Transaction,
	}
}
