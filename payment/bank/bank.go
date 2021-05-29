package bank

import (
	"fmt"
	"time"
)

func ProcessPayment(fromAccount int, toAccount int, amount int) error {
	fmt.Printf("Transfered %d  from %d to %d at %v via bank transfer", amount, fromAccount, toAccount, time.Now().String())
	return nil
}
