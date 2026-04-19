package commission

import "math"

func Calculate(alifCard bool, amount int) int {
	if alifCard {
		return 0
	}
	commissionFloat := float64(amount) * 0.0029
	return int(math.Ceil(commissionFloat))
}
