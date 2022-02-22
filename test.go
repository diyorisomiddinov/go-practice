package main

func main() {

	num := [5]int{3, 5, 7, 8, 9}
	var check int
	for i, v := range num {
		check = num[i]
		for j, k := range num {
			if check == k {
				delete(num, j)
			}
		}
	}

}
