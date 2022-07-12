package bitwise

/*
1-10:

00000001
00000010
00000011
00000100
00000101

00000110
00000111
00001000
00001001
00001010

for every leftshift, the number is squared
this means 1, 2, 4, 8, ... have the same number of 1 bits
use this for memoization

*/

func countBits(n int) []int {
    ans := make([]int, n + 1)
    
    for i := 1 ; i <= n ; i++ {
        // the (i&1) is needed for odd numbers
        // even numbers will always have 0 as last bit
        // odd numbers will always have 1 as last bit
        ans[i] = ans[i / 2] + (i & 1)
    }
    
    return ans
}