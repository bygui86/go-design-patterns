package adapters

import (
	"math/rand"

	"github.com/bygui86/go-design-patterns/payment/bank"
)

type BankAdapter struct{}

func (b *BankAdapter) Send(fromEmail string, toEmail string, amount int) error {
	fromAccount := findAccountByEmail(fromEmail)
	toAccount := findAccountByEmail(toEmail)
	return bank.ProcessPayment(fromAccount, toAccount, amount)
}

func findAccountByEmail(email string) int {
	// This will query the data store to return the Account Number
	return rand.Int()
}
