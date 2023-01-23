package account

import (
	"root/model/currency"
	"root/model/person"
)

type Account struct {
	User     person.Person       `json:"User"`
	Currency []currency.Currency `json:"Currency"`
}

// Empty array
func Empty() *[]Account {
	return &[]Account{}
}

//Empty single object
// func Empty() *Account {
// 	return &Account{}
// }

func New(Person person.Person, Currency []currency.Currency) *Account {
	return &Account{
		Person, Currency,
	}
}
