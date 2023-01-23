package person

type Person struct {
	Username  string `json:"Username"`
	AccountID string `json:"AccountID"`
}

func New(Username, AccountID string) *Person {
	return &Person{
		Username, AccountID,
	}
}
