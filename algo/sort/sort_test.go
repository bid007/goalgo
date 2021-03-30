package sort

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{11, 10, 9, 0, -1}
	InsertionSort(arr)
	assert.ElementsMatch(t, arr, []int{-1, 0, 9, 10, 11})

	arr = []int{1, 2, 3, 4, 5, 6, 7}
	assert.ElementsMatch(t, arr, []int{1, 2, 3, 4, 5, 6, 7})

	arr = []int{10, 9, 8, 7, 3, 2, 0, -100}
	assert.ElementsMatch(t, arr, []int{-100, 0, 2, 3, 7, 8, 9, 10})
}

func TestMergeSort(t *testing.T) {
	assert.ElementsMatch(t, MergeSort([]int{11, 10, 9, 0, -1}), []int{-1, 0, 9, 10, 11})
	assert.ElementsMatch(t, MergeSort([]int{1, 2, 3, 4, 5, 6, 7}), []int{1, 2, 3, 4, 5, 6, 7})
	assert.ElementsMatch(t, MergeSort([]int{10, 9, 8, 7, 3, 2, 0, -100}), []int{-100, 0, 2, 3, 7, 8, 9, 10})
}

func BenchmarkInsertionSort(b *testing.B) {
	arr := make([]int, b.N)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		arr[i] = rand.Intn(b.N)
	}
	InsertionSort(arr)
}

func BenchmarkMergeSort(b *testing.B) {
	arr := make([]int, b.N)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		arr[i] = rand.Intn(b.N)
	}
	MergeSort(arr)
}
