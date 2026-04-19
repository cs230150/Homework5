package receipt

import (
	"fmt"
	"math/rand"
	"time"
)

func TransactionId() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}
func Card(card string) string {
	last4 := card[len(card)-4:]
	return "UZCARD**" + last4
}
func CurrentTime() string {
	return time.Now().Format("02.01.2006 15:04:05")
}
func FullName(name, surname string) string {
	return fmt.Sprintf("%s %s", surname, name)
}
func FormatedAmount(amount int) string {
	a := fmt.Sprintf("%d", amount)
	b := len(a)

	var result string
	for i, c := range a {
		result += string(c)
		if (b-i-1)%3 == 0 && i != b-1 {
			result += " "
		}
	}
	return result
}
