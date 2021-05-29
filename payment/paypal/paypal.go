package paypal

import (
	"fmt"
)

func Send(senderMobile, receiverMobile string, amount int) error {
	fmt.Printf("Send %d  from %s to %s via paypal transfer", amount, senderMobile, receiverMobile)
	return nil
}
