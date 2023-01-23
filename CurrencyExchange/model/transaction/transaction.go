package transaction

type Transaction struct {
	TransactionID uint   `json:"TransactionID"`
	From          string `json:"From"`
	To            string `json:"To"`
	Deposit       uint   `json:"Deposit"`
}

func New(TransactionID uint, From, To string, Deposit uint) *Transaction {
	return &Transaction{
		TransactionID, From, To, Deposit,
	}
}
