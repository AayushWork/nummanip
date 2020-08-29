package calc

import "errors"

//Add function adds 2 int and returns an int type
func Add(numbers ...int) (error, int) {
	sum := 0

	if len(numbers) < 2 {
		return errors.New("Provide more than 2 numbers"), sum
	} else {
		for _, num := range numbers {
			sum = sum + num
		}

		return nil, sum
	}
}
