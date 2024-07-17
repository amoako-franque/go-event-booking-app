package main

import "strings"

func userInputValidator(firstname string, lastname string, email string, tickerQty uint) (bool, bool, bool) {
	isUsernameValid := len(firstname) >= 2 && len(lastname) >= 2
	isEmailValid := strings.Contains(email, "@")
	isValidTicketNUmber := tickerQty > 0 && tickerQty <= ticketsLeft

	return isUsernameValid, isEmailValid, isValidTicketNUmber
}
