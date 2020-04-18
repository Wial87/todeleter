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
		"4eb5b280b761f9c7289901d109554bb2001f6bf6ce4caf35dcb4f4de06e4ea843b548e4a6d60e74c09d40": "дд",
		"b3ddcccff151dc93030646f454d5dba17d559ef3bb724512f51223c9965f267208f4c3f1dadc85f8f73ed": "дд", // Это пример, поля не должны быть пустыми!!!
	}

	// Функция запуска аккаунтов
	bot.StartAccounts(accounts)
}
