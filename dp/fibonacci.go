package dp

func fibo(n int) int {
	dict := make(map[int]int)

	// no need to pass pointer; maps are always passed by reference
	return faster(n, dict)

}

func faster(n int, dict map[int]int) {
	if n == 0 { return 0 }
	if n == 1 { return 1 }

	if dict[n] > 0 { return dict[n] }

	dict[n] = faster(n - 1, dict) + faster(n - 2, dict)
	return dict[n]
}