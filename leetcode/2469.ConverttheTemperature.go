package leet

func convertTemperature(celsius float64) []float64 {
	answer := []float64{celsius + 273.15, celsius*1.8 + 32}
	return answer
}
