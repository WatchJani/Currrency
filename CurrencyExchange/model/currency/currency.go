package currency

type Currency struct {
	CurrencyID uint    `json:"CurrencyID"`
	Currency   string  `json:"Currency"`
	Amount     float64 `json:"Amount"`
}

func New(currency string, CurrencyID uint, Amount float64) *Currency {
	return &Currency{
		CurrencyID, currency, Amount,
	}
}
