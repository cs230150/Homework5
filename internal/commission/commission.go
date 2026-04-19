package commission

import (
	"errors"
	"math"
)

func ValidateAmount(amount int) error {
	if amount < 500 {
		return errors.New("ошибка: сумма меньше 500 сум")
	}
	if amount > 15000000 {
		return errors.New("ошибка: сумма больше 15 000 000 сум")
	}
	return nil
}

func CalculateCommission(amount int, isAlif bool) int {
	if isAlif {
		return 0 // Alif карта - без комиссии
	}
	return int(math.Ceil(float64(amount) * 0.0029))
}
