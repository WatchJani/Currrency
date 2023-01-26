package currency

func FromTo(Currency *[]Currency, Find string, To string, Deposit uint) {
	//zasto pointer ne radi na *Currency
	for index := range *Currency {
		if (*Currency)[index].Currency == Find {
			(*Currency)[index].Amount -= float64(Deposit)
			continue
		}
		if (*Currency)[index].Currency == To {
			(*Currency)[index].Amount += (float64(Deposit) * 0.5)
			continue
		}
	}
}
