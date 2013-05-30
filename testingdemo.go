package gotestingdemo

func NextEvenNotMultipleByFive(i int) int {
	if i%2 == 1 {
		i++
	} else {
		i = i + 2
	}
	return i
}
