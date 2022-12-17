package main

import (
	"fmt"
	"sort"
)

// Returns the k most frequent elements in the given slice
func kMostFrequent(nums []int, k int) []int {
	// Return an empty slice if the input slice is empty
	if len(nums) == 0 {
		return []int{}
	}

	// Return the entire slice if k is greater than the number of elements
	if k > len(nums) {
		return nums
	}

	// Use a map to count the frequency of each element
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	// Create a min-heap of elements and their frequencies, sorted by frequency
	heap := make([][]int, 0, len(count))
	for num, freq := range count {
		heap = append(heap, []int{num, freq})
	}
	sort.Slice(heap, func(i, j int) bool { return heap[i][1] < heap[j][1] })

	// Pop the top k elements off the heap
	result := make([]int, 0, k)
	for len(result) < k {
		result = append(result, heap[0][0])
		heap = heap[1:]
	}

	// Sort the resulting slice by frequency in descending order
	sort.Slice(result, func(i, j int) bool { return count[result[i]] > count[result[j]] })

	return result
}

func main() {
	// Unit tests
	tests := []struct {
		nums   []int
		k      int
		expect []int
	}{
		// Test empty slice
		{
			nums:   []int{},
			k:      5,
			expect: []int{},
		},
		// Test slice with fewer elements than k
		{
			nums:   []int{1, 2, 3},
			k:      5,
			expect: []int{1, 2, 3},
		},
		// Test slice with more elements than k
		{
			nums:   []int{1, 2, 2, 3, 3, 3},
			k:      2,
			expect: []int{3, 2},
		},
		// Test slice with all elements the same
		{
			nums:   []int{1, 1, 1, 1, 1},
			k:      3,
			expect: []int{1},
		},
	}

	for _, test := range tests {
		result := kMostFrequent(test.nums, test.k)
		if fmt.Sprint(result) == fmt.Sprint(test.expect) {
			fmt.Println("PASS")
		} else {
			fmt.Println("FAIL")
		}
	}
}
