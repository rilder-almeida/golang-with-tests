package arrayslices

func SumArray(array []int) int {
	sum := 0
	for _, value := range array {
		sum += value
	}
	return sum
}

func SumAll(arrays ...[]int) int {
	var sums []int
	for _, array := range arrays {
		sums = append(sums, SumArray(array))
	}
	return SumArray(sums)
}

func SumTails(arrays ...[]int) int {
	var sums []int
	for _, array := range arrays {
		if len(array) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, SumArray(array[1:]))
		}
	}
	return SumArray(sums)
}
