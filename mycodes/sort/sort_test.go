package sort_test

import (
	"fmt"
	"math/rand"
	"time"
	"testing"
	"mycodes/sort"
)

func TestSort(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	num := rand.Intn(100)
	slice := make([]int, num)
	backup := make([]int, num)
	for i:=0; i<num; i++ {
		backup[i] = rand.Intn(1000)
	}
	fmt.Printf("%s\n", "original squence")
	fmt.Printf("%v\n", backup)

	copy(slice, backup)
	fmt.Printf("1.bubble sort:\n")
	sort.BubbleSort(slice)
	fmt.Printf("%v\n", slice)

	copy(slice, backup)
	fmt.Printf("2.selection sort:\n")
	sort.SelectionSort(slice)
	fmt.Printf("%v\n", slice)

	copy(slice, backup)
	fmt.Printf("3.insertion sort:\n")
	sort.InsertionSort(slice)
	fmt.Printf("%v\n", slice)

	copy(slice, backup)
	fmt.Printf("4.binary insertion sort:\n")
	sort.BinaryInsertionSort(slice)
	fmt.Printf("%v\n", slice)

	copy(slice, backup)
	fmt.Printf("5.shell sort:\n")
	sort.ShellSort(slice)
	fmt.Printf("%v\n", slice)

	copy(slice, backup)
	fmt.Printf("6.merge sort:\n")
	sort.MergeSort(slice)
	fmt.Printf("%v\n", slice)

	copy(slice, backup)
	fmt.Printf("7.quick sort:\n")
	sort.QuickSort(slice)
	fmt.Printf("%v\n", slice)

	copy(slice, backup)
	fmt.Printf("8.heap sort:\n")
	sort.HeapSort(slice)
	fmt.Printf("%v\n", slice)

	copy(slice, backup)
	fmt.Printf("9.counting sort:\n")
	sort.CountingSort(slice)
	fmt.Printf("%v\n", slice)

	copy(slice, backup)
	fmt.Printf("10.bucket sort:\n")
	sort.BucketSort(slice)
	fmt.Printf("%v\n", slice)

	copy(slice, backup)
	fmt.Printf("11.radix sort:\n")
	sort.RadixSort(slice)
	fmt.Printf("%v\n", slice)
}