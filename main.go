package main

import (
	"deleter/bot"
)

func main() {
	/*
		Вот пример как вбивать аккаунты:
			accounts := map[string]string{
			"token": "keyWord",
			"token": "keyWord",
		}
	*/

	accounts := map[string]string{
		// "token": "keyWord",
		"06ce2c2d15f25f052b642e58325982ce85153fb5a9cbb511531e2db4303f8311a71f823cf78482fc8fc8f": "дд",
		"d6fd8d4781e7cab77a5e538534f71b4ec258874157063af52c73ce161717f419cd55d4b8d8c544a99a0cc": "дд", // Это пример, поля не должны быть пустыми!!!
	}

	// Функция запуска аккаунтов
	bot.StartAccounts(accounts)
}
