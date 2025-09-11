package main

func main() {
	int_slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, val := range int_slice {
		if val % 2 == 0 {
			println(val, " is even")
		} else {
			println(val, " is odd")
		}
	}
}