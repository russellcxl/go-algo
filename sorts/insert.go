package sorts

/*

time: O(nlogn)
space: O(n) or, for worst case, O(n^2)

1. take pivot as the last
2. loop through all the rest and keep an index j
3. compare num[i] to pivot; if smaller, swap with j, j++
4. at the end, swap num[j] with pivot

*/


func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}