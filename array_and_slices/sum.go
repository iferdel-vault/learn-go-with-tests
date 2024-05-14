package main

func Sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func SumAllTails(slicesOfNumbers ...[]int) (sliceOfTailSum []int) {
	for _, slice := range slicesOfNumbers {
        if len(slice) == 0 {
            sliceOfTailSum = append(sliceOfTailSum, 0)
            continue
        }
        tail := slice[1:]
		sliceOfTailSum = append(sliceOfTailSum, Sum(tail))
	}
	return sliceOfTailSum
}
