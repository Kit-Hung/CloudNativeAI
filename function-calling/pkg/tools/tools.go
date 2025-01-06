package tools

func AddTool(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func SubTool(numbers []int) int {
	result := 0
	for i, number := range numbers {
		if i == 0 {
			result = number
		} else {
			result -= number
		}
	}
	return result
}
