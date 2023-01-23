package currency

type Currency struct {
	CurrencyID uint   `json:"CurrencyID"`
	Currency   string `json:"Currency"`
	Amount     uint   `json:"Amount"`
}

func New(currency string, CurrencyID, Amount uint) *Currency {
	return &Currency{
		CurrencyID, currency, Amount,
	}
}
