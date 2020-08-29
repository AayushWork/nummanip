package calc

//Add function adds 2 int and returns an int type
func Add(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum = sum + num
	}

	return sum
}
