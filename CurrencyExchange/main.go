package main

import (
	"root/model/account"
	"root/model/transfer"
	"root/os"
)

func main() {
	accounts := account.Empty()
	os.ReadFile(accounts, "./data/Account.json")

	transfers := transfer.Empty()
	os.ReadFile(transfers, "./data/Transaction.json")

	os.WriteFile(transfers)
}
