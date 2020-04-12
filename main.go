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
		//"3439ba6cce7d15993687e092daf191f5b9b465f81c1741d2f4f67b6c65c04538f6b80ac6daf1a22ff33af": "дд", // Это пример, поля не должны быть пустыми!!!
	}

	// Функция запуска аккаунтов
	bot.StartAccounts(accounts)
}
