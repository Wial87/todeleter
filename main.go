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
		"токен": "триггер-слово",
		"токен": "триггер-слово", // Это пример, поля не должны быть пустыми!!!
	}

	// Функция запуска аккаунтов
	bot.StartAccounts(accounts)
}
