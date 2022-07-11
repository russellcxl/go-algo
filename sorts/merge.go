package sorts

/*

time: O(nlogn)
space: O(n) because of extra array needed to sort subarrays

1. keep dividing the arrays until there's just 1 element left
2. merge the elements, sort them and return to the previous callback
3. the previous callback will receive 2 sorted arrays, sort them, and return up the chain

*/

func mergeSort(items []int) {

	if len(items) < 2 { return items }

	left := mergeSort(items[:len(items) / 2])
	right := mergeSort(items[len(items) / 2:])

	return merge(left, right)
}

func merge(a, b []int) []int {
	final := make([]int, len(a) + len(b))

	var i, j int
    for i < len(a) && j < len(b) {
        if a[i] < b[j] {
            final = append(final, a[i])
            i++
        } else {
            final = append(final, b[j])
            j++
        }
    }
    for ; i < len(a); i++ {
        final = append(final, a[i])
    }
    for ; j < len(b); j++ {
        final = append(final, b[j])
    }

    return final
}