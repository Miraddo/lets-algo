package main

// based on : https://leetcode.com/problems/reverse-integer/

func Reverse(x int) int {
	var rev int

	for x != 0 {
		pop := x % 10
		x /= 10

		rev = rev*10 + pop

		if !(rev > -2147483648 && rev < 2147483647) {
			return 0
		}
	}
	return rev

}



func main(){

}
