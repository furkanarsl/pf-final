package utils

func CalculatePercent(price float64, percent int16) float64 {
	return (price * (float64(percent) / 100))
}
