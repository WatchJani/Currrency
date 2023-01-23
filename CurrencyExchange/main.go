package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"root/model/account"
)

func main() {
	account := account.Empty()

	file, err := os.Open("./data/Account.json")

	if err != nil {
		log.Fatalln(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(account)
	fmt.Printf("%s\n", data)

	fmt.Println(account)
	err = json.Unmarshal(data, &account)



	fmt.Println(err)
	fmt.Println((*account)[0].Currency[0].Amount)
}
