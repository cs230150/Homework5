package commission

import (
	"errors"
	"math"
)

func Calculate(amount int, isAlif bool) (int, error) {
	if amount < 500 || amount > 15000000 {
		return 0, errors.New("сумма вне диапазона")
	}

	if isAlif {
		return 0, nil
	}

	commissionFloat := float64(amount) * 0.0029
	return int(math.Ceil(commissionFloat)), nil
}
