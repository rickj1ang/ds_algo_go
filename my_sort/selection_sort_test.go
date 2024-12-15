package mysort

import "testing"

// it is a very rough test U should ues fuzzing to randam test the function
// I will do it latter
func TestSelectionSort(t *testing.T) {
	// declare a slice
	nums := []int{12, 42, 5, 23, 50, 22, 8, 19}

	// call the selectionSort()
	selectionSort(nums)

	//a for loop check the result
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > nums[i+1] {
			t.Error("sort wrong")
		}
	}
}
