package binary

func binarySearch(target int, numbers []int) int {
	lastIndex := len(numbers) - 1
	firstIndex := 0
	for firstIndex <= lastIndex {
		middleIndex := (firstIndex + lastIndex) / 2
		result := numbers[middleIndex]
		if result == target {
			return middleIndex
		}
		if target < result {
			lastIndex = middleIndex - 1
		} else {
			firstIndex = middleIndex + 1
		}
	}
	return -1
}
