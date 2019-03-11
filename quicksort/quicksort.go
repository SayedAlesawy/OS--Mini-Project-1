package quicksort

import (
	"math/rand"
)

//getRandomPartitionIndex A function to get a random partition index
func getRandomPartitionIndex(arr []int, l int, r int) int{
	//Get a random partition index
	var randomIdx int = (rand.Int() % (r-l)) + l

	//Swap the random idx with the last index in the array
	arr[randomIdx], arr[r-1] = arr[r-1], arr[randomIdx]

	var partitionIdx int = l
	var pivot int = arr[r-1]

	for j := l; j < r - 1; j++{
		//If an element is smaller than or equal to the pivot
		//Move it to the left half
		if arr[j] <= pivot{
			arr[partitionIdx], arr[j] = arr[j], arr[partitionIdx]
			partitionIdx = partitionIdx + 1
		}
	}

	//Place the pivot in its correct position
	arr[partitionIdx], arr[r-1] = arr[r-1], arr[partitionIdx]

	return partitionIdx
}

//Sort a function to apply the quick sort algorithm
func Sort(arr []int, l int, r int){
	if r - l <= 1{
		return
	}

	var p int = getRandomPartitionIndex(arr, l, r)

	Sort(arr, l, p)
	Sort(arr, p, r)
}