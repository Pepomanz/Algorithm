package binary

import "testing"

func TestFindBinarySearch(t *testing.T) {
	numbers := []int{1, 5, 6, 9, 10, 30, 33, 45, 56, 78}
	given := 30
	expect := 5
	actual := binarySearch(given, numbers)
	if expect != actual {
		t.Errorf("Index Not Found")
	}
}
